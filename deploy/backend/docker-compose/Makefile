# *******************************************************************************
# MAKEFILE_TYPE :=
# Makefile for {{ GITHUB_PROJECT }}/{{ PROJECT_TYPE }}
# - [](<https://www.gnu.org/software/make/manual/make.html>)
# - [](<https://www.gnu.org/software/make/manual/html_node/Simple-Makefile.html>)
# *******************************************************************************
# Purpose:
#	This script is used to build, test, and deploy the project.
# *******************************************************************************
# Usage:
#	make [target]
# *******************************************************************************
# History:
#	2021/09/01	Clark Hsu  First release
# *******************************************************************************
# *******************************************************************************
# Variables
# TOP_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
TOP_DIR := $(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
TARGETS := $(shell ls -p scripts | grep -v /)
GIT_PROVIDER := {{ GIT_PROVIDER }}
PORJECGT_USER := {{ PORJECGT_USER }}
GITHUB_PROJECT := gin-realtime
# CONTAINER_NAME := {{ GITHUB_PROJECT }}

CONTAINER_NAME := backend
CONTAINER_HOST := 0.0.0.0
CONTAINER_PORT := 3001

# *******************************************************************************
# *******************************************************************************
# Functions
# *******************************************************************************
# *******************************************************************************
# Main
# *******************************************************************************
# INTERNAL VARIABLES
# Read all subsequent tasks as arguments of the first task
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(args) $(RUN_ARGS):;@:)
# *******************************************************************************
.DEFAULT_GOAL := help

.PHONY: help
help:  ## Help
	@echo "Usage: make [target]"
	@echo "Targets:"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf " \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all:  ## all
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo -e "\n>>> ...\n"
	@# make clean
	make build
	@# make status
	@# make version
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: $(TARGETS)
$(TARGETS):	 ## Run the function
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo ">>> targets ..."
	@ls -p scripts | grep -v / | xargs -I {} chmod +x scripts/{}
	@./scripts/$@
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: func_list
func_list:	## List all functions
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo ">>> targets ..."
	@for TARGET in $(TARGETS); do \
		echo $$TARGET; \
	done
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

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
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: nocache
nocache:  ## nocache
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo -e "\n>>> nocache...\n"
	$(TOP_DIR)/$(GITHUB_PROJECT)
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

.PHONY: logs
logs:  ## logs
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo -e "\n>>> logs...\n"
	@# make clean
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: status
status:	 ## Status the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Status the project"
	@echo ">>> $(TOP_DIR)..."
	@echo ">>> $(GITHUB_USER)/$(GITHUB_PROJECT)/$(PACKAGE_VERSION)..."
	@echo ">>> $(DISTRO)/$(OS)/$(ARCH)..."
	@echo
	docker-compose ps
	@echo
	docker-compose images
	@# docker ps -a
	@# docker images
	@echo
	docker volume ls
	@echo
	docker network ls
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
	docker-compose -f docker-compose.yaml up -d backend
	@echo "Backend: http://0.0.0.0:3001"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: undeploy
undeploy:  ## Undeploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Undeploy the project"
	docker-compose -f docker-compose.yaml stop backend
	@docker-compose -f docker-compose.yaml rm -f backend
	@docker volume rm docker-compose_backend_data
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

# *******************************************************************************
# EOF
# *******************************************************************************
