# Patipada Search System

This project is a part of the Waris Lakthong senior project, focusing on developing a hybrid search system for Dhammanva question-answer videos. The system is designed using a microservices architecture, with each backend service adhering to clean architecture principles.

## Services Overview

1. [Frontend Service](./frontend/)
2. [Search Service](./search-esdb-service/)
3. [Authentication Service](./auth-service/)
4. [Data Management Service](./data-management-service/)

## Data Source

All data is extracted from Dhammanva live question-answer video transcripts ([visit the youtube channel](https://www.youtube.com/@dhammanava7327)), which include:

- Start time of the section
- End time of the section
- Question
- Answer
- YouTube URL

The data is stored in the [search-esdb-service/data](./search-esdb-service/data/) directory.

## Deployment

In [.github/workflows](./.github/workflows/) directiory contians workflow for each container and each workflow will run on `workflow_dispatch`, means you need to press a `run` button in github actions page.
for more information please visit [.github/workflows](./.github/workflows/)

## Run Locally

### Prerequisites

1. docker install
2. Make install (optional)

### Steps

- Clone the project

- Copy `.env.template`, paste in the same directory then rename to `.env` (for more information in env variable navigate to [Environment Variables](#environment-variables))

- Open your terminal and type

```bash
make up_build
```

or

```bash
docker compose -f docker-compose.dev.yml up -d
```

## Environment Variables

<a id="environment-variables"></a>

To run this project, you will need to add the following environment variables to your .env file

`ELASTIC_USERNAME` : username to access elasticSearch

`ELASTIC_PASSWORD`: password to access elasticSearch

`AUTH_DB_USER` : username to access Auth DB (PostgreSQL)

`AUTH_DB_PASSWORD` : password to access Auth DB (PostgreSQL)

`JWT_KEY` : key used to sign JWT token

`SENDER_EMAIL` : email of sender use to send email in reset password process

`SENDER_PASSWORD` : app password (Gmail) of sender

`LINK_URL` : url link to provide reset password token (local would be frontend, production would be host name)

`SUPER_ADMIN_PASSWORD` : Password for the super admin account

`SUPER_ADMIN_EMAIL` : Email for the super admin account

`ADMIN_PASSWORD` : Password for the admin account

`ADMIN_EMAIL` : Email for the admin account

`USER_PASSWORD` : Password for the user account

`USER_EMAIL` : Email for the user account

`DATA_MNGMNT_DB_USER` : Default username for data management database

`DATA_MNGMNT_DB_PASSWORD` : password for default user in data management database

> _<u>Note</u>_ if you run each project without docker e.g. `go run main.go` you do not need to assign the .env variables each service has default env variables in app.env except `SENDER_PASSWORD`, which you are required to assign in `app.env`

## Tech Stack

**Frontend:** React, Typescript, Vite, Chakra UI, Tiptap

**Seasrch service:** Golang, Gin, ElasticSearch

**Authentication service:** Golang, Echo, PostgreSQL

**Data management service:** Golang, Gin, MongoDB

**Communication:** 

**Containerization:** Docker

## Project structure

###

    .
    ├── auth-db                     # initialize authentication database
    ├── auth-service
    ├── data-management-service
    ├── elastic                     # initialize elasticSearch
    ├── frontend
    ├── .github/workflows           # Workflow files
    ├── nginx                       # example of nginx used in production
    ├── search-esdb-service
    ├── .env.template               # template of .env
    ├── .secrets.template           # template of .secrets used in act (test workflow locally)
    ├── README.md
    ├── Makefile                    # Makefile used in development
    ├── docker-compose.dev.yml      # used to build image and test local
    ├── docker-compose.prod.yml     # used in production

> You can discover more information for each directory in `README.md` of each directory

> This project is under development
