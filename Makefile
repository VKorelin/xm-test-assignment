build: generate-proto
	cd company && GOOS=linux GOARCH=amd64 make build

precommit: generate-proto
	cd company && make precommit

generate-proto:
	cd company && make generate

sqlc-generate:
	cd company && sqlc generate

init-db:
	docker compose up -d --build --remove-orphans company_db

run: 
	docker-compose up -d --force-recreate --remove-orphans --build

migrate:
	cd company && make migrate

start: init-db build run migrate

stop:
	docker compose down

#UTILS

# 1. add new migration
MIGRATION_NAME:=""

add-company-migration:
	cd company/migrations && goose create $(MIGRATION_NAME) sql