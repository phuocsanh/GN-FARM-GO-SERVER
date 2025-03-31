GOOSE_DBSTRING ?= "root:123456@tcp(127.0.0.1:3308)/GO_ECOMMERCE"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql


# name app
APP_NAME = server

# run:
# 	go run ./cmd/$(APP_NAME)/

# build:
# 	go build -o bin/migration_app ./cmd/$(APP_NAME)/

# dev:
# 	 docker-compose up && go run ./cmd/$(APP_NAME)

# kill:
# 	docker-compose kill

# up:
# 	docker-compose up -d

# down:
# 	docker-compose down

start:
	 docker compose up && go run ./cmd/$(APP_NAME)

docker_stop:
	docker compose down

docker_up:
	docker compose up -d

docker_build:
	docker compose up -d --build

dev:
	go run ./cmd/$(APP_NAME)
up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
# Create new migration
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) reset

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: start dev downse upse resetse docker_build docker_stop docker_up swag