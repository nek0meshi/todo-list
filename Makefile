BACKEND_CONTAINER_NAME := backend
DB_CONTAINER_NAME := db
PROD_PROJECT_NAME := todo-list-prod

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: down-v
down-v:
	docker compose down -v

.PHONY: build
build:
	docker compose build

.PHONY: b-sh
b-sh:
	docker compose exec ${BACKEND_CONTAINER_NAME} bash

.PHONY: d-sh
d-sh:
	docker compose exec ${DB_CONTAINER_NAME} bash

.PHONY: yarn
yarn:
	cd frontend && yarn && yarn dev

.PHONY: go-fmt
go-fmt:
	docker compose exec ${BACKEND_CONTAINER_NAME} go fmt .

.PHONY: db-setup
db-setup:
	docker compose exec ${DB_CONTAINER_NAME} sh -c 'exec mysql -u sample -psample sample < ./sql/setup.sql'

.PHONY: prod-up
prod-up:
	docker compose -p ${PROD_PROJECT_NAME} -f docker-compose-prod.yml up

.PHONY: prod-build
prod-build:
	docker compose -p ${PROD_PROJECT_NAME} -f docker-compose-prod.yml build

.PHONY: prod-down
prod-down:
	docker compose -p ${PROD_PROJECT_NAME} down

.PHONY: prod-db-setup
prod-db-setup:
	docker compose -p ${PROD_PROJECT_NAME} exec ${DB_CONTAINER_NAME} sh -c 'exec mysql -u sample -psample sample < ./sql/setup.sql'

.PHONY: prod-f-sh
prod-f-sh:
	docker compose -p ${PROD_PROJECT_NAME} exec ${FRONTEND_CONTAINER_NAME} sh

.PHONY: create-project
create-project:
	docker run \
		--rm \
		--volume $(PWD):/work:cached \
		--workdir /work \
		node:14-alpine \
		yarn create vite-app frontend
