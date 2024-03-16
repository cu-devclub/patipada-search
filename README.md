# Patipada Search System

This project is a part of the Waris Lakthong senior project, focusing on developing a hybrid search system for Dhammanva question-answer videos. The system is designed using a microservices architecture, with each backend service adhering to clean architecture principles.

## !! Important Note

To deploy / develop each service or each database please check README.md for each one. This README.md provided only overall of the project.

## Services Overview

1. [Frontend Service](./frontend/)
2. [Search Service](./search-esdb-service/)
3. [Authentication Service](./auth-service/)
4. [Data Management Service](./data-management-service/)
5. [Machine learning Service](./ml-service/)

## Data Source

### Record

All data is extracted from Dhammanva live question-answer video transcripts ([visit the youtube channel](https://www.youtube.com/@dhammanava7327)), which include:

- Start time of the section
- End time of the section
- Question
- Answer
- YouTube URL

The data is stored in the [data](./data/record/) directory.

### Stop words

The stop word dictionary use in this project is currently from pythainlp
as store in [stopwords](./data/stopword/)

## Deployment

In [.github/workflows](./.github/workflows/) directiory contians workflow for each container and each workflow will run on `workflow_dispatch`, means you need to press a `run` button in github actions page.
for more information please visit [.github/workflows](./.github/workflows/)

## Development Process (Run locally)

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

- Or in case you want to run each service individually, you can navigate to each service and reading through README.md

## Production

By default you should use github workflow and deploy to produciton by trigger each jobs you want. However if you want something more you can use docker-compose.prod.yml and Makefile to set the service as you want.
| These 2 files already have ci action to savely deploy to server

- docker-compose.prod.yml ; manipulate through docker compose

- Makefile is the collection of commands I prebuilt which you can run using `make <command>` all the details for each command provided in the comment section in Makefile

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

`SUPER_ADMIN_USERNAME` : Username for the super admin account

`SUPER_ADMIN_PASSWORD` : Password for the super admin account

`SUPER_ADMIN_EMAIL` : Email for the super admin account

`ADMIN_PASSWORD` : Password for the admin account

`ADMIN_EMAIL` : Email for the admin account

`USER_PASSWORD` : Password for the user account

`USER_EMAIL` : Email for the user account

`DATA_MNGMNT_DB_USER` : Default username for data management database

`DATA_MNGMNT_DB_PASSWORD` : password for default user in data management database

`DOCKERHUB_USERNAME` : username for docker hub use only in production

`RABBITMQ_USERNAME` : username for rabbit mq service; this value will be set as default user and needed services will used this value as credential

`RABBITMQ_PASSWORD` : password for rabbit mq service; this value will be set as default user and needed services will used this value as credential

`GRAFANA_USERNAME` : username for login to grafana dashboard

`GRAFANA_PASSWORD` : password for login to grafana dashboard

> _<u>Note</u>_ if you run each project without docker e.g. `go run main.go` you do not need to assign the .env variables each service has default env variables in app.env except `SENDER_PASSWORD`, which you are required to assign in `app.env`

## Tech Stack

**Frontend:** React, Typescript, Vite, Chakra UI, Tiptap

**Search service:** Golang, ElasticSearch

**Authentication service:** Golang, PostgreSQL

**Data management service:** Golang, MongoDB

**Machine Learning service:** Python, Flask

**Communication:** GRPC, RabbitMQ

**Containerization:** Docker

**Monitoring:** Loki, Promtail, Grafana

## Project structure

###

    .
    ├── auth-db                     # initialize authentication database
    ├── auth-service
    ├── data-management-service
    ├── elastic                     # initialize elasticSearch
    ├── frontend
    ├── .github/workflows           # Workflow files (CI/CD)
    ├── nginx                       # nginx file for both dev and prod
    ├── search-esdb-service
    ├── ml-service
    ├── data                        # Store data source
        ├── record                  # Records (Q&A , start & end time, URL, ...)
        ├── stopword                # List of stopword
    |- monitoring                   # directory store config file for monitoring related
    ├── .env.template               # template of .env
    ├── .secrets.template           # template of .secrets used in act (test workflow locally)
    ├── README.md
    ├── Makefile                    # Makefile used in development
    ├── Makefile.prod               # Makefile used in production
    ├── docker-compose.dev.yml      # used to build image and test local
    ├── docker-compose.prod.yml     # used in production

> You can discover more information for each directory in `README.md` of each directory

> This project is under development
