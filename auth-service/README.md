# Authentication service 
> authentication service for dhammanva search system

This service use to authenticate and authorize user.

#### CI/CD : 
[deploy auth service](../.github/workflows/auth-deploy.yml)
[deploy auth db](../.github/workflows/auth-db.yml)

## Run locally 
If you do not want to run using docker you can run using golang

#### Prerequisite 

1. Install golang 
2. Start auth-db docker compose
``` bash
cd <your-path>/<project-root(where docker compose is)>
```
``` bash
docker compose -f docker-compose.dev.yml up -d auth-db
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
      docker compose -f docker-compose.dev.yml up -d auth-service 
      ```

## API Reference
Postman documentation : [click](https://documenter.getpostman.com/view/14178897/2s9YsFFaVj)

## Tech stack 
**Language** : Golang
**Web Server** : Echo
**Database connection**: GORM + PostGreSQL

## Project structure
This service imply [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
### 

    .
    ├── auth_proto              # gRPC server      
    ├── config            
    ├── database                    
    ├── errors                  # Custom errors
    ├── jwt                     # JWT token related
    ├── messages                # Response Message
    ├── server                     
    ├── tests                   # unit & integration testing  
    ├── users
        ├── handlers           
        ├── entities
        ├── usecases
        ├── repositories
        ├── models
        ├── migration           # migrate default user
        ├── helper              # helper function  
    ├── app.env                 # default env file
    ├── go.mod               
    ├── go.sum               
    ├── main.go              
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
