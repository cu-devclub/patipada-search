# Data Management service

> data management service for dhammanva search system

This service use to store managed data

#### CI/CD : 
[deploy data service](../.github/workflows/data-deploy.yml)
[deploy data db](../.github/workflows/data-db.yml)

## Run locally

#### Prerequisite

1. Install golang
2. Install docker / make
3. Start data management db
``` bash
cd <your-path>/<project-root(where docker compose is)>
```
``` bash
docker compose -f docker-compose.dev.yml up -d data-db
```
4. Copy [data folder](../data/) to this directory and rename to `datasource`

#### Steps
There are 3 ways to run the service 
  1. Using golang (Recommend for isolate service development)
      ``` bash
      cd <your-path>/data-managment-service
      ```
      ```bash
      go mod tidy
      go mod vendor
      go run mock/isolate.go 
      ```
  2. Using docker
    - uncomment every line in [Dockerfile](./Dockerfile)
    - Navigate to root directory
    - Run
      ```bash
      docker compose -f docker-compose.dev.yml up --build -d data-service 
      ```
  3. Using make (spin up all dependencies service) (recommend for final testing / testing with another services)
    - Install make
    - run
     ``` bash
      cd <your-path>/patipada-search
      ```
      ```bash
      make up_build_data
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
    ├── mock                    # mock external services
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
