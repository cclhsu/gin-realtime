#*******************************************************************************
# MAKEFILE_TYPE :=
# Makefile for {{ projectName }}/{{ projectType }}
#*******************************************************************************
# Purpose:
#	This script is used to build, test, and deploy the project.
#*******************************************************************************
# Usage:
#	make [target]
#*******************************************************************************
# History:
#	2021/09/01	Clark Hsu  First release
#*******************************************************************************
#*******************************************************************************
# Variables
TOP_DIR=$(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
GIT_PROVIDER := {{ gitProvider }}
PORJECGT_USER := {{ projectUser }}
PROJECT_NAME := {{ projectName }}
DOCKER_FILE := Dockerfile
DOCKER_REGISTRY := {{ dockerRegistry }}
DOCKER_PROJECT := gin-realtime
# latest | mac | wintel | linux | $(shell git rev-parse --short HEAD)
DOCKER_TAG := latest
DOCKER_PORT := 3001
DOCKER_USER := {{ dockerUser }}
DOCKER_PASSWORD := {{ dockerPassword }}

#*******************************************************************************
#*******************************************************************************
# Functions
#*******************************************************************************
#*******************************************************************************
# Main
#*******************************************************************************
# INTERNAL VARIABLES
# Read all subsequent tasks as arguments of the first task
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(args) $(RUN_ARGS):;@:)
#*******************************************************************************
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf " \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init:  ## Initialize the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Initialize the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: install
install:  ## Install packages for the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Install packages for the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: update
update:	 ## Update packages for the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Update packages for the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: build
build:	## Build the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Build the project"
	@# docker buildx install
	@# docker buildx create --name mybuilder
	@# docker buildx ls
	@# docker buildx use mybuilder
	docker buildx use default
	docker buildx inspect --bootstrap
	docker build --rm -f $(DOCKER_FILE) -t $(DOCKER_PROJECT):$(DOCKER_TAG) --no-cache --load ../../../
	@# docker buildx build --rm -f $(DOCKER_FILE) --platform linux/amd64 -t $(DOCKER_PROJECT):$(DOCKER_TAG) --no-cache --load ../../../
	@# docker buildx build --rm -f $(DOCKER_FILE) --platform linux/amd64,windows/amd64 -t $(DOCKER_PROJECT):$(DOCKER_TAG) --no-cache --load ../../../
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: start
start:	## Start the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Start the project"
	docker run --name $(DOCKER_PROJECT) -p $(DOCKER_PORT):$(DOCKER_PORT) -v '$(shell pwd)/../../../.env:/app/.env' --network docker-compose_my-network -d $(DOCKER_PROJECT):$(DOCKER_TAG)
	@echo "$(DOCKER_PROJECT): http://0.0.0.0:$(DOCKER_PORT)"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: stop
stop:  ## Stop the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Stop the project"
	docker stop $(DOCKER_PROJECT) 2>/dev/null; true
	docker rm $(DOCKER_PROJECT) 2>/dev/null; true
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: bash
bash:  ## Bash the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Bash the project"
	# docker exec --rm -it --name $(DOCKER_PROJECT) $(DOCKER_PROJECT):$(DOCKER_TAG) bash
	docker run --rm -it --name $(DOCKER_PROJECT) $(DOCKER_PROJECT):$(DOCKER_TAG) sh
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: status
status:	 ## Status the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Status the project"
	docker ps -a
	docker images
	docker volume ls
	docker network ls

	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: log
log:  ## Log the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Log the project"
	docker logs $(DOCKER_PROJECT)
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: test
test:  ## Test the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Test the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: lint
lint:  ## Lint the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Lint the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: package
package:  ## Package the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Package the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: deploy
deploy:	 ## Deploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Deploy the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: undeploy
undeploy:  ## Undeploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Undeploy the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: clean
clean:	## Clean the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Clean the project"
	@# docker system prune --force --all --volumes 2>/dev/null; true
	@# docker buildx prune --force --all --volumes 2>/dev/null; true
	docker rm $(shell docker ps -a -q) 2>/dev/null; true
	docker rmi $(shell docker images -f "dangling=true" -q) 2>/dev/null; true
	docker volume prune --force 2>/dev/null; true
	docker network prune --force 2>/dev/null; true
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: publish-image
publish-image:	## publish docker image
	@echo "Publishing docker image..."
	@# docker login -u $(DOCKER_USER) -p $(DOCKER_PASWORD)
	docker tag $(DOCKER_PROJECT):$(DOCKER_TAG) $(DOCKER_USER)/$(DOCKER_PROJECT)
	docker push $(DOCKER_USER)/$(DOCKER_PROJECT)

.PHONY: untag-publish-image
untag-publish-image:  ## untag and publish docker image
	@echo "Untag and publishing docker image..."
	docker rmi $(DOCKER_USER)/$(DOCKER_PROJECT):$(DOCKER_TAG)

#*******************************************************************************
# EOF
#*******************************************************************************
