build: generate-proto
	cd company && GOOS=linux GOARCH=amd64 make build

precommit: generate-proto
	cd company && make precommit

generate-proto:
	cd company && make generate

sqlc-generate:
	cd company && sqlc generate

run: 
	docker-compose up -d --force-recreate --remove-orphans --build

migrate:
	cd company && make migrate

#UTILS

# 1. add new migration
MIGRATION_NAME:=""

add-company-migration:
	cd company/migrations && goose create $(MIGRATION_NAME) sql