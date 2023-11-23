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
GITHUB_USER := {{ GITHUB_USER }}
GITHUB_PROJECT := {{ GITHUB_PROJECT }}
NAMESPACE := "{{ GITHUB_PROJECT }}"
LABEL := {{ GITHUB_PROJECT }}
CONTEXT := {{ CONTEXT }}

NAMESPACE := "default"
LABEL := "io.kompose.service={{ GITHUB_PROJECT }}"
POD_NAME := $(shell kubectl get pods -l ${LABEL} -n $(NAMESPACE) -o jsonpath='{range .items[*]}{.metadata.name}')
CONTAINER_HOST := 0.0.0.0
CONTAINER_PORT := {{ SERVICE_PORT }}

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
	@echo "Backend: http://0.0.0.0:3000"
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

# *******************************************************************************
# EOF
# *******************************************************************************
