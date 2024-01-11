# Data Management service

> data management service for dhammanva search system

This service use to store managed data

## Run locally

If you do not want to run using docker you can run using golang

#### Prerequisite

1. Install golang

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

Postman documentation : // Not available yet

## Tech stack

**Language** : Golang
**Web Server** : Gin
**Database**: Mongo DB

## Project structure

This service imply [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

###

    .
    ├── auth_proto              # gRPC with auth service
    ├── config
    ├── database
    ├── errors                  # Custom errors
    ├── messages                # Response Message
    ├── server
    ├── tests                   # integration testing
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
