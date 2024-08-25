CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=xm/company/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

# gRPC generation

LOCAL_BIN:=$(CURDIR)/bin
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc
COMPANY_PROTO_OUT_PATH:="pkg/api/company/v1"
COMPANY_PROTO_PATH:="api/company/v1"

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest

.vendor-proto: .vendor-rm vendor-proto/validate 

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

PHONY: .protoc-generate
.protoc-generate:
	mkdir -p ${COMPANY_PROTO_OUT_PATH}
	$(PROTOC) -I ${COMPANY_PROTO_PATH} -I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out ${COMPANY_PROTO_OUT_PATH} --go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out ${COMPANY_PROTO_OUT_PATH} --go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:${COMPANY_PROTO_OUT_PATH}" \
	${COMPANY_PROTO_PATH}/company.proto
	go mod tidy

PHONY: generate
generate: .bin-deps .vendor-proto .protoc-generate

#migrations:
.PHONY: migrate
migrate:
	goose -dir migrations postgres postgresql://postgres:password@localhost:5433/company up