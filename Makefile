build-all: generate-proto build-projects

build-all-fast: generate-proto-fast build-projects

build-projects:
	cd company && GOOS=linux GOARCH=amd64 make build

run-containers:
	docker-compose up -d --force-recreate --remove-orphans --build

precommit: generate-proto
	cd company && make precommit

generate-proto:
	cd company && make generate

generate-proto-fast:
	cd company && make fast-generate

run-all: build-all run-containers

run-all-fast: build-all-fast run-containers