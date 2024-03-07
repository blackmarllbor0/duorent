APP_NAME = duorent
APP_PATH = cmd/$(APP_NAME).go
ENV_MODE ?= dev

.PHONY:run
run:
	go run $(APP_PATH)


.PHONY:build
build:
	docker-compose --env-file .env.$(ENV_MODE) \
		--profile tools run --rm postgres-migrate up

.PHONY:lint
lint:
	golangci-lint run