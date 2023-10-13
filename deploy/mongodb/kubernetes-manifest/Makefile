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
# GITHUB_PROJECT := {{ githubProject }}
GIT_COMMIT:=$(shell git rev-parse --short HEAD 2>/dev/null)-$(shell date "+%Y%m%d%H%M%S")
GIT_TAG:=$(shell git describe --tags --dirty 2>/dev/null)
# NAMESPACE := "{{ githubProject }}"
# LABEL := {{ githubProject }}
CONTEXT := {{ context }}

NAMESPACE := "default"
LABEL := "io.kompose.service=mongodb"
POD_NAME := $(shell kubectl get pods -l ${LABEL} -n $(NAMESPACE) -o jsonpath='{range .items[*]}{.metadata.name}')
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
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: status
status:	 ## Status the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Status the project"
	# kubectl get all -o wide --show-labels # -A
	@echo ">>> Getting kubectl pod status..."
	kubectl get pods -o wide --show-labels # --all-namespaces
	@echo
	@echo ">>> Getting kubectl service status..."
	kubectl get services -o wide --show-labels # --all-namespaces
	@echo
	@echo ">>> Getting kubectl deployment status..."
	kubectl get deployments -o wide --show-labels # --all-namespaces

	@# echo ">>> Getting kubectl ingress status..."
	@# kubectl get ingress -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl daemonset status..."
	@# kubectl get daemonset -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl statefulset status..."
	@# kubectl get statefulset -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl job status..."
	@# kubectl get job -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl cronjob status..."
	@# kubectl get cronjob -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl hpa status..."
	@# kubectl get hpa -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl ingressclass status..."
	@# kubectl get ingressclass -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl configmap status..."
	@# kubectl get configmap -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl secret status..."
	@# kubectl get secret -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl persistentvolumeclaim status..."
	@# kubectl get pvc -o wide --show-labels # --all-namespaces
	@# @echo ">>> Getting kubectl persistentvolume status..."
	@# kubectl get pv -o wide --show-labels # --all-namespaces

	@echo
	@echo ">>> Expose Ports for kubectl services..."
	@# kubectl get services --output custom-columns=NAME:.metadata.name,PORT:.spec.ports[*].nodePort
	kubectl get services --output custom-columns=NAME:.metadata.name,PORT:.spec.ports[*].targetPort
	@echo
	@echo ">>> Getting redis-cli status..."
	kubectl exec -it ${POD_NAME} -- mongosh --host -u your_db_user -p your_db_pass --eval "db.stats()"
	kubectl exec -it ${POD_NAME} -- mongosh --host -u your_db_user -p your_db_pass --eval "db.serverStatus()"
	kubectl exec -it ${POD_NAME} -- mongosh --host -u your_db_user -p your_db_pass --eval "db.version()"
	kubectl exec -it ${POD_NAME} -- mongosh --host -u your_db_user -p your_db_pass -eval "db.hostInfo()"

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
	kubectl logs -l app=$(LABEL)
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: kubectl-events
kubectl-events:	 ## kubectl get events
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Getting kubectl events..."
	kubectl get events --sort-by=.metadata.creationTimestamp
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
	kubectl apply -f .
	@echo "MongoDB: http://${CONTAINER_HOST}:${CONTAINER_PORT}"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: undeploy
undeploy:  ## Undeploy the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Undeploy the project"
	kubectl delete -f .
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: kubectl-context
kubectl-context:  ## kubectl config get-contexts
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Getting kubectl context..."
	kubectl config get-contexts
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: kubectl-set-context
kubectl-set-context:  ## kubectl config use-context $(CONTEXT)
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Setting kubectl context..."
	kubectl config use-context $(CONTEXT)
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: clean
clean:	## Clean the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Clean the project"
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

#*******************************************************************************
# EOF
#*******************************************************************************
