FRONTEND_CONTAINER_NAME := frontend
BACKEND_CONTAINER_NAME := backend

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: down-v
down-v:
	docker-compose down -v

.PHONY: f-sh
f-sh:
	docker-compose exec ${FRONTEND_CONTAINER_NAME} sh

.PHONY: b-sh
b-sh:
	docker-compose exec ${BACKEND_CONTAINER_NAME} sh

.PHONY: yarn
yarn:
	docker-compose exec ${FRONTEND_CONTAINER_NAME} sh -c "yarn && yarn dev"

.PHONY: run-hello-world
run-hello-world:
	docker-compose exec ${BACKEND_CONTAINER_NAME} sh -c "go install hello_world && ../bin/hello_world"

.PHONY: create-project
create-project:
	docker run \
		--rm \
		--volume $(PWD):/work:cached \
		--workdir /work \
		node:14-alpine \
		yarn create vite-app frontend
