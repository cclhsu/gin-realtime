#*******************************************************************************
# MAKEFILE_TYPE :=
# Makefile for {{ githubProject }}/{{ projectType }}
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
DISTRO ?= $(shell cat /etc/*-release 2>/dev/null | uniq -u | grep ^ID= | cut -d = -f 2 | sed s/\"//g | sed s/linux/-linux/g && sw_vers -productName 2>/dev/null | sed 's/ //g' | tr A-Z a-z)
OS ?= $(shell uname -s | tr A-Z a-z)
ARCH ?= $(shell uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')
GIT_PROVIDER := {{ gitProvider }}
GITHUB_USER := {{ githubUser }}
GITHUB_PROJECT := {{ githubProject }}
GIT_COMMIT:=$(shell git rev-parse --short HEAD 2>/dev/null)-$(shell date "+%Y%m%d%H%M%S")
GIT_TAG:=$(shell git describe --tags --dirty 2>/dev/null)
# CONTAINER_NAME := {{ githubProject }}

CONTAINER_NAME := mongodb
CONTAINER_HOST := 0.0.0.0
CONTAINER_PORT := 27017

ifneq ("$(wildcard .env)", "")
	include .env
endif

ifneq ("$(wildcard VERSION.txt)", "")
	PACKAGE_VERSION=$(shell grep -i PACKAGE_VERSION VERSION.txt | cut -d '=' -f 2 | tr -d '[:space:]')
else ifdef LATEST
	PACKAGE_VERSION=latest
else
	PACKAGE_VERSION ?= $(shell curl -s "https://api.github.com/repos/${GITHUB_USER}/${GITHUB_PROJECT}/releases/latest" | jq --raw-output .tag_name)
endif
TAG=$(PACKAGE_VERSION)

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
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
	brew tap mongodb/brew
	brew install mongodb-community
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
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: start
start:	## Start the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Start the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: stop
stop:  ## Stop the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Stop the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: bash
bash:  ## Bash the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Bash the project"
	@# docker-compose exec ${CONTAINER_NAME} bash
	docker-compose exec ${CONTAINER_NAME} sh
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: status
status:	 ## Status the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Status the project"
	docker-compose ps
	docker-compose images
	@# docker ps -a
	@# docker images
	docker volume ls
	docker network ls

	@# @echo
	@# @echo "Status the project with docker-compose"
	@# docker-compose exec ${CONTAINER_NAME} mongosh --host -u your_db_user -p your_db_pass --eval "db.stats()"
	@# docker-compose exec ${CONTAINER_NAME} mongosh --host -u your_db_user -p your_db_pass --eval "db.serverStatus()"
	@# docker-compose exec ${CONTAINER_NAME} mongosh --host -u your_db_user -p your_db_pass --eval "db.version()"
	@# docker-compose exec ${CONTAINER_NAME} mongosh --host -u your_db_user -p your_db_pass -eval "db.hostInfo()"

	@echo
	@echo "Status the project with mongosh through port ${CONTAINER_PORT}"
	mongosh --host ${CONTAINER_HOST} --port ${CONTAINER_PORT} -u your_db_user -p your_db_pass --eval "db.stats()"
	@# mongosh --host ${CONTAINER_HOST} --port ${CONTAINER_PORT} -u your_db_user -p your_db_pass --eval "db.serverStatus()"
	@# mongosh --host ${CONTAINER_HOST} --port ${CONTAINER_PORT} -u your_db_user -p your_db_pass --eval "db.version()"
	@# mongosh --host ${CONTAINER_HOST} --port ${CONTAINER_PORT} -u your_db_user -p your_db_pass -eval "db.hostInfo()"

	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: log
log:  ## Log the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Log the project"
	docker logs ${CONTAINER_NAME}
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
	docker-compose -f docker-compose.yaml up -d ${CONTAINER_NAME}
	@echo "MongoDB: http://${CONTAINER_HOST}:${CONTAINER_PORT}"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: undeploy
undeploy:  ## Undeploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Undeploy the project"
	docker-compose -f docker-compose.yaml stop ${CONTAINER_NAME}
	@docker-compose -f docker-compose.yaml rm -f ${CONTAINER_NAME}
	@docker volume rm docker-compose_mongodb_data
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

#*******************************************************************************
# EOF
#*******************************************************************************
