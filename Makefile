build-all: generate-proto build-projects

build-all-fast: generate-proto-fast build-projects

build-projects: sqlc-generate
	cd company && GOOS=linux GOARCH=amd64 make build

run-containers:
	docker-compose up -d --force-recreate --remove-orphans --build

precommit: generate-proto
	cd company && make precommit

generate-proto:
	cd company && make generate

generate-proto-fast:
	cd company && make fast-generate

migrate:
	cd company && make migrate

sqlc-generate:
	cd company && sqlc generate

run-all: build-all run-containers migrate

run-all-fast: build-all-fast run-containers migrate

#UTILS

# 1. add new migration
MIGRATION_NAME:=""

add-company-migration:
	cd company/migrations && goose create $(MIGRATION_NAME) sql