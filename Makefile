AUTH_BINARY = authApp
SEARCH_BINARY = searchApp
DATA_BINARY = dataApp
## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose -f docker-compose.dev.yml up -d
	@echo "Docker images started!"

## up_build: stops docker compose (if running), builds all projects and starts docker compose
up_build:  build_auth build_search build_data build_frontend 
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d auth-service search-service data-service frontend nginx rabbitmq
	@echo "Docker images built and started!"

up_build_service_monitoring:  build_auth build_search build_data build_frontend 
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d auth-service search-service data-service frontend nginx rabbitmq loki promtail grafana
	@echo "Docker images built and started!"

up_build_backend: build_auth build_search build_data
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d auth-service search-service data-service rabbitmq
	@echo "Docker images built and started!"

####### AUTH SERVICE #######
## up_build_auth: stops docker compose (if running), builds projects and starts docker compose
up_build_auth: build_auth
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down auth-service
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d auth-service
	@echo "Docker images built and started!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} .
	@echo "Done!"

## down_auth: stops the auth service
down_auth:
	@echo "Stopping auth service..."
	docker compose -f docker-compose.dev.yml down auth-service
	@echo "Auth service stopped!"

## up_dev_auth: stops db container and rebuild and start go server
up_dev_auth:
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down auth-db
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d auth-db
	@echo "Docker images built and started!"
	cd auth-service && go run main.go
	@echo "Auth service development server started!"
#################################

###### Search Service ######
## up_build_search: stops docker compose (if running), builds projects and starts docker compose
up_build_search: build_search
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down search-service
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d search-service
	@echo "Docker images built and started!"

## build_search: builds the search binary as a linux executable
build_search:
	@echo "Building search binary..."
	cd search-esdb-service && env GOOS=linux CGO_ENABLED=0 go build -o ${SEARCH_BINARY} .
	@echo "Done!"

## down_search: stops the search service
down_search:
	@echo "Stopping search service..."
	docker compose -f docker-compose.dev.yml down search-service
	@echo "Search service stopped!"

## empty_elastic: stops elastic-db (if running), remove volumes, starts elastic-db
## !!! BE CAREFUL WITH THIS COMMAND CUZ IT WILL REMOVE ALL THE EXISITING DATA!!!
empty_elastic:
	@echo "Stopping elastic-db (if running...)"
	docker compose -f docker-compose.dev.yml down elastic-db
	@echo "Remove volumes..."
	rm -rf ./volumes/database/elastic
	@echo "starting elastic-db containers..."
	docker compose -f docker-compose.dev.yml up --build -d elastic-db
	@echo "elastic-db started!"

## up_dev_search: stops db container and rebuild and start go server
up_dev_search:
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down elastic-db
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d elastic-db
	@echo "Docker images built and started!"
	cd search-esdb-service && go run main.go
	@echo "Search service development server started!"

#################################

###### Data Management Service ######
## up_build_data: stops docker compose (if running), builds projects and starts docker compose
up_build_data: build_data
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down data-service
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d data-service
	@echo "Docker images built and started!"

## build_data: builds the data binary as a linux executable
build_data:
	@echo "Building data binary..."
	cd data-management-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DATA_BINARY} .
	@echo "Done!"


## down_data: stops the data service
down_data:
	@echo "Stopping data service..."
	docker compose -f docker-compose.dev.yml down data-service
	@echo "Data service stopped!"

## empty_data_db: stops data-db (if running), removes volumes and starts data-db
## !!! BE CAREFUL WITH THIS COMMAND CUZ IT WILL REMOVE ALL THE EXISITING DATA!!!
empty_data_db:
	@echo "Stopping data-db (if running...)"
	docker compose -f docker-compose.dev.yml down data-db
	@echo "Remove volumes..."
	rm -rf ./volumes/database/mongo-data
	@echo "starting data-db containers..."
	docker compose -f docker-compose.dev.yml up --build -d data-db
	@echo "Data-db started!"

## up_dev_data: stops db container and rebuild and start go server
up_dev_data:
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down data-db
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d data-db
	@echo "Docker images built and started!"
	cd data-management-service && go run main.go
	@echo "Data service development server started!"

#################################

##### Frontend Service #####
## up_build_frontend: stops docker compose (if running), builds projects and starts docker compose
up_build_frontend: build_frontend
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down frontend
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d frontend
	@echo "Docker images built and started!"

## down_frontend: stops the frontend service
down_frontend:
	@echo "Stopping frontend service..."
	docker compose -f docker-compose.dev.yml down frontend
	@echo "Frontend service stopped!"

## build_frontend: builds the frontend binary as a linux executable
build_frontend:
	@echo "Building frontend binary..."
	cd frontend && yarn build
	@echo "Done!"

#################################

###### Nginx Service ######
## up_build_nginx: stops docker compose (if running), builds projects and starts docker compose
up_build_nginx:
	@echo "Stopping docker images (if running...)"
	docker compose -f docker-compose.dev.yml down nginx
	@echo "Building (when required) and starting docker images..."
	docker compose -f docker-compose.dev.yml up --build -d nginx
	@echo "Docker images built and started!"



## down_nginx : stops the nginx service
down_nginx:
	@echo "Stopping nginx service..."
	docker compose -f docker-compose.dev.yml down nginx
	@echo "Nginx service stopped!"

## down: stops all containers
down:
	@echo "Stopping Docker images..."
	docker compose -f docker-compose.dev.yml down
	@echo "Docker images stopped!"


## monitoring service
up_monitoring:
	@echo "Starting monitoring service..."
	docker compose -f docker-compose.dev.yml up -d loki promtail grafana prometheus
	@echo "Monitoring service started!"

down_monitoring:
	@echo "Stopping monitoring service..."
	docker compose -f docker-compose.dev.yml down loki promtail grafana prometheus
	@echo "Monitoring service stopped!"