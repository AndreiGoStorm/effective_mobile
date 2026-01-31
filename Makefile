BIN := "./bin/subscription"

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/application

install-lint:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin

swag-init:
	swag init -g cmd/application/main.go -o cmd/application/docs

up-build:
	cd docker && \
	docker compose --project-name="subscription" up --build

up:
	cd docker && \
	docker compose --project-name="subscription" up -d

down:
	cd docker && \
	docker compose --project-name="subscription" stop
