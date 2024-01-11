# Authentication service 
> authentication service for dhammanva search system

This service use to authenticate and authorize user.

## Run locally 
If you do not want to run using docker you can run using golang

#### Prerequisite 

1. Install golang 

#### Steps
1. Open your terminal

2. Naviage to this directory 
```bash
cd <your-path>/auth-service
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
Postman documentation : [click](https://documenter.getpostman.com/view/14178897/2s9YsFFaVj)

## Tech stack 
**Language** : Golang
**Web Server** : Echo
**Database connection**: GORM + PostGreSQL

## Project structure
This service imply [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
### 

    .
    ├── auth_proto              # gRPC with auth service      
    ├── config            
    ├── database                    
    ├── errors                  # Custom errors
    ├── jwt                     # JWT token related
    ├── messages                # Response Message
    ├── server                     
    ├── tests                   # integration testing  
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
