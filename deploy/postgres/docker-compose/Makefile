#*******************************************************************************
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

CONTAINER_NAME := postgres
CONTAINER_HOST := 0.0.0.0
CONTAINER_PORT := 5432

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
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf " \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init:  ## Initialize the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Initialize the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: install
install:  ## Install packages for the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Install packages for the project"
	@# brew search postgresql
	brew install postgresql@16
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
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT version();"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT current_database();"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT current_user;"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT * FROM pg_stat_activity;"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT * FROM pg_stat_activity WHERE state <> 'idle';"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "CREATE EXTENSION IF NOT EXISTS pg_stat_statements;"
	@# docker-compose exec ${CONTAINER_NAME} psql -U your_db_user -d your_db_name -c "SELECT query, calls, total_time, rows FROM pg_stat_statements ORDER BY calls DESC LIMIT 10;"

	@echo
	@echo "Status the project with psql through port ${CONTAINER_PORT}"
	psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT version();"
	psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT current_database();"
	psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT current_user;"
	psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';"
	@# psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT * FROM pg_stat_activity;"
	@# psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT * FROM pg_stat_activity WHERE state <> 'idle';"
	@# psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "CREATE EXTENSION IF NOT EXISTS pg_stat_statements;"
	@# psql postgres://your_db_user:your_db_pass@${CONTAINER_HOST}:${CONTAINER_PORT}/your_db_name -c "SELECT query, calls, total_time, rows FROM pg_stat_statements ORDER BY calls DESC LIMIT 10;"
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
	@echo "Postgres: http://${CONTAINER_HOST}:${CONTAINER_PORT}"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: undeploy
undeploy:  ## Undeploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Undeploy the project"
	docker-compose -f docker-compose.yaml stop ${CONTAINER_NAME}
	@docker-compose -f docker-compose.yaml rm -f ${CONTAINER_NAME}
	@docker volume rm docker-compose_postgres_data
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
