.DEFAULT_GOAL := help

DOCKER_REPO ?= quay.io
DOCKER_PROFILE ?= openbazaar
DOCKER_NAME ?= status-server
DOCKER_TAG ?= $(shell git describe --tags --abbrev=0)
DOCKER_IMAGE_NAME ?= $(DOCKER_REPO)/$(DOCKER_PROFILE)/$(DOCKER_NAME):$(DOCKER_TAG)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docker: ## Create Docker image
	docker build -t $(DOCKER_IMAGE_NAME) .

docker_push: ## Push Docker image to registry
	docker push $(DOCKER_IMAGE_NAME)