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

1. Open your terminal

2. Naviage to this directory

```bash
cd <your-path>/data-management-service
```

3. Run

```bash
go get ./...
```

```bash
go mod vendor
```

```bash
go run main.go
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
