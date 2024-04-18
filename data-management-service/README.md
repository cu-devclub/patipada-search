# Data Management service

> data management service for dhammanva search system

This service use to store managed data

#### CI/CD : 
[deploy data service](../.github/workflows/data-deploy.yml)
[deploy data db](../.github/workflows/data-db.yml)

## Run locally

If you do not want to run using docker you can run using golang

#### Prerequisite

1. Install golang
2. Start data management db
``` bash
cd <your-path>/<project-root(where docker compose is)>
```
``` bash
docker compose -f docker-compose.dev.yml up -d data-db
```

#### Steps
There are 2 ways to run the service 
  1. Using golang 
      ``` bash
      cd <your-path>/search-esdb-service
      ```
      ```bash
      go get ./...
      go mod vendor
      go run main.go 
      ```
  2. Using docker
    - uncomment every line in [Dockerfile](./Dockerfile)
    - Navigate to root directory
    - Run
      ```bash
      docker compose -f docker-compose.dev.yml up -d data-service 
      ```


## API Reference

Postman documentation : [link](https://documenter.getpostman.com/view/14178897/2s9YsNdVnb)

## Tech stack

**Language** : Golang
**Web Server** : Gin
**Database**: Mongo DB

## Project structure

This service imply [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

###

    .
    ├── communication           # gRPC client setup
    ├── proto                   # GRPC
        ├── auth_proto 
        ├── search_proto 
    ├── config
    ├── database
    ├── errors                  # Custom errors
    ├── messages                # Response Message
    ├── server
    ├── tests                   # unit test & integration testing
    ├── structValidator         # validator
    ├── request
        ├── handlers
        ├── entities
        ├── usecases
        ├── repositories
        ├── models
        ├── helper
    ├── app.env                 # default env file
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
