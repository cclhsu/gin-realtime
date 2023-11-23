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
TOP_DIR := $(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
GIT_PROVIDER := github.com
PORJECGT_USER := cclhsu
GITHUB_PROJECT := gin-realtime

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
	export GO111MODULE=on
	go mod init ${GIT_PROVIDER}/${GITHUB_USER}/${GITHUB_PROJECT}
	go get github.com/asaskevich/govalidator
	go get github.com/gin-gonic/gin
	go get github.com/go-redis/redis/v8
	go get github.com/golang-jwt/jwt
	go get github.com/google/uuid
	go get github.com/gorilla/websocket
	go get github.com/joho/godotenv
	go get github.com/patrickmn/go-cache
	go get github.com/sirupsen/logrus
	go get github.com/swaggo/files
	go get github.com/swaggo/gin-swagger
	go get github.com/swaggo/swag
	go get go.mongodb.org/mongo-driver/bson
	go get go.mongodb.org/mongo-driver/mongo
	go get go.mongodb.org/mongo-driver/mongo/options
	go get go.mongodb.org/mongo-driver/mongo/readpref
	go get google.golang.org/grpc
	go get google.golang.org/grpc/codes
	go get google.golang.org/grpc/metadata
	go get google.golang.org/grpc/reflection
	go get google.golang.org/grpc/status
	go get google.golang.org/protobuf/reflect/protoreflect
	go get google.golang.org/protobuf/runtime/protoimpl
	go get google.golang.org/protobuf/types/known/emptypb
	go get gorm.io/driver/postgres
	go get gorm.io/gorm
	go mod tidy
	go mod vendor
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: update
update:	 ## Update packages for the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Update packages for the project"
	export GO111MODULE=on
	go get -u
	go mod tidy
	go mod vendor
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: build
build:	## Build the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Build the project"
	@# swag init -g cmd/${GITHUB_PROJECT}/main.go -o doc/openapi
	@# go build -o ./bin/${GITHUB_PROJECT} ./cmd/${GITHUB_PROJECT}

	@echo
	@SERVICE_NAME=elasticsearch-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	SERVICE_NAME=elasticsearch-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=graphql-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=graphql-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=grpc-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	SERVICE_NAME=grpc-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=kafka-consumer-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=kafka-producer-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=redis-consumer-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=redis-producer-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=server-sent-event-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=server-sent-event-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=socket-io-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=socket-io-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webhook-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webhook-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webpush-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webpush-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webrtc-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=webrtc-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=websocket-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME
	@echo
	@SERVICE_NAME=websocket-server-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi; \
	go build -o ./bin/$$SERVICE_NAME ./cmd/$$SERVICE_NAME

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
	swag init -g cmd/${GITHUB_PROJECT}/main.go -o doc/openapi
	go build -o ./bin/${GITHUB_PROJECT} ./cmd/${GITHUB_PROJECT}
	psql postgres://your_db_user:your_db_pass@0.0.0.0:5432/your_db_name -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";" || :
	go run cmd/${GITHUB_PROJECT}/main.go migrate
	./bin/${GITHUB_PROJECT}
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
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: test
test:  ## Test the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Test the project"
	swag init -g cmd/${GITHUB_PROJECT}/main.go -o doc/openapi
	go build -o ./bin/${GITHUB_PROJECT} ./cmd/${GITHUB_PROJECT}
	go test -v
	go test -cover
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: lint
lint:  ## Lint the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Lint the project"
	swag init -g cmd/${GITHUB_PROJECT}/main.go -o doc/openapi
	go build -o ./bin/${GITHUB_PROJECT} ./cmd/${GITHUB_PROJECT}
	golangci-lint run
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
	go clean -i -r -cache -testcache -modcache
	rm -rf ${TOP_DIR}/${GITHUB_PROJECT}
	rm -rf ${TOP_DIR}/data/bin
	rm -rf ${TOP_DIR}/{bin,dist,target,vendor}
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: generate_openapi
generate_openapi:  ## Generate openapi for the project
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	# @echo "Generate openapi for the project"
	# swag init -g cmd/${GITHUB_PROJECT}/main.go -o doc/openapi

	rm -rf doc/openapi/* && ls -al doc/openapi
	# @echo
	# @SERVICE_NAME=elasticsearch-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=elasticsearch-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	@echo
	@SERVICE_NAME=graphql-client-service; \
	swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=graphql-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=grpc-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=grpc-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=kafka-consumer-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=kafka-producer-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=redis-consumer-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=redis-producer-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=server-sent-event-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=server-sent-event-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=socket-io-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=socket-io-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webhook-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webhook-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webpush-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webpush-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webrtc-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=webrtc-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=websocket-client-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi
	# @echo
	# @SERVICE_NAME=websocket-server-service; \
	# swag init -g cmd/$$SERVICE_NAME/main.go -o doc/openapi

	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_graphql_server
run_graphql_server:	 ## Run the graphql server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the graphql server"
	./bin/graphql-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_graphql_client
run_graphql_client:	## Run the graphql client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the graphql client"
	cd ${TOP_DIR}/client/graphql && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_graphql_client1
run_graphql_client1:  ## Run the graphql client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the graphql client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/graphql-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_graphql_client2
run_graphql_client2:  ## Run the graphql client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the graphql client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/graphql-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_grpc_server
run_grpc_server:  ## Run the grpc server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the grpc server"
	./bin/grpc-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_grpc_client
run_grpc_client:	## Run the grpc client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the grpc client"
	cd ${TOP_DIR}/client/grpc && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_grpc_client1
run_grpc_client1:  ## Run the grpc client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the grpc client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/grpc-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_grpc_client2
run_grpc_client2:  ## Run the grpc client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the grpc client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/grpc-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_kafka_producer
run_kafka_producer:	 ## Run the kafka producer
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the kafka producer"
	./bin/kafka-producer-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_kafka_client
run_kafka_client:	## Run the kafka client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the kafka client"
	cd ${TOP_DIR}/client/kafka && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_kafka_consumer
run_kafka_consumer:	 ## Run the kafka consumer
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the kafka consumer"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/kafka-consumer-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_redis_producer
run_redis_producer:	 ## Run the redis producer
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the redis producer"
	./bin/redis-producer-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_redis_client
run_redis_client:	## Run the redis client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the redis client"
	cd ${TOP_DIR}/client/redis && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_redis_consumer
run_redis_consumer:	 ## Run the redis consumer
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the redis consumer"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/redis-consumer-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_server_sent_event_server
run_server_sent_event_server:  ## Run the server sent event server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event server"
	./bin/server-sent-event-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_server_sent_event_client
run_server_sent_event_client:	## Run the server_sent_event client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server_sent_event client"
	cd ${TOP_DIR}/client/server_sent_event && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_server_sent_event_client1
run_server_sent_event_client1:	## Run the server sent event client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/server-sent-event-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_server_sent_event_client2
run_server_sent_event_client2:	## Run the server sent event client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/server-sent-event-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_socket_io_server
run_socket_io_server:  ## Run the server sent event server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event server"
	./bin/socket-io-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_socket_io_client
run_socket_io_client:	## Run the socket_io client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the socket_io client"
	cd ${TOP_DIR}/client/socket_io && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_socket_io_client1
run_socket_io_client1:	## Run the server sent event client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/socket-io-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_socket_io_client2
run_socket_io_client2:	## Run the server sent event client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the server sent event client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/socket-io-client-service

.PHONY: run_webhook_server
run_webhook_server:	 ## Run the webhook server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webhook server"
	./bin/webhook-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webhook_client
run_webhook_client:	## Run the webhook client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webhook client"
	cd ${TOP_DIR}/client/webhook && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webhook_client1
run_webhook_client1:  ## Run the webhook client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webhook client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/webhook-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webhook_client2
run_webhook_client2:  ## Run the webhook client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webhook client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/webhook-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webpush_server
run_webpush_server:	 ## Run the webpush server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webpush server"
	./bin/webpush-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webpush_client
run_webpush_client:	## Run the webpush client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webpush client"
	cd ${TOP_DIR}/client/webpush && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webpush_client1
run_webpush_client1:  ## Run the webpush client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webpush client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/webpush-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webpush_client2
run_webpush_client2:  ## Run the webpush client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webpush client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/webpush-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webrtc_server
run_webrtc_server:	 ## Run the webrtc server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webrtc server"
	./bin/webrtc-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_webrtc_client
run_webrtc_client:	## Run the webrtc client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the webrtc client"
	cd ${TOP_DIR}/client/webrtc && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_websocket_server
run_websocket_server:  ## Run the websocket server
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the websocket server"
	./bin/websocket-server-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_websocket_client
run_websocket_client:	## Run the websocket client
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the websocket client"
	cd ${TOP_DIR}/client/websocket && http-server -p 3003
	@# cd ${TOP_DIR}/client/websocket && python3 -m http.server --cgi 3003
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_websocket_client1
run_websocket_client1:	## Run the websocket client1
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the websocket client1"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3002 ./bin/websocket-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

.PHONY: run_websocket_client2
run_websocket_client2:	## Run the websocket client2
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ..."
	@echo "Run the websocket client2"
	SERVER_HOST=0.0.0.0 SERVER_PORT=3001 SERVICE_PORT=3003 ./bin/websocket-client-service
	@echo ">>> [$$(date +'%Y-%m-%d %H:%M:%S')] $@ ... Done"

# *******************************************************************************
# EOF
# *******************************************************************************
