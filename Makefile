# APP ?= github.com/marcelotoledo5000/game-log
APP ?= game-log
GOLANG_VERSION ?= 1.20
VERSION ?= $(shell git describe --tags --abbrev=0)
DOCKER_IMAGE ?= skygvinn/$(APP):$(VERSION)
DOCKER_DEFAULT_PARAMS ?= --rm -v "$(PWD):/go/src/$(APP)" -w "/go/src/$(APP)" golang:$(GOLANG_VERSION)

.DEFAULT_GOAL = package

.PHONY: package
package:
	docker build --pull --no-cache -t $(DOCKER_IMAGE) --build-arg GOLANG_VERSION=$(GOLANG_VERSION) .

.PHONY: release
release: package
	docker push $(DOCKER_IMAGE)

.PHONY: test
test:
	docker run $(DOCKER_DEFAULT_PARAMS) go test ./...

.PHONY: test_local
test_local:
	go test ./...

.PHONY: build_local
build_local:
	go build

.PHONY: run
run:
	go run main.go $(ARGS)
