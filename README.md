# Patipada Search System

This project is a part of the Waris Lakthong senior project, focusing on the development of a hybrid search system for Dhammanva question-answer videos. The system is designed using a microservices architecture, with each backend service adhering to clean architecture principles.

## Services Overview

1. **Frontend Service**
2. **Search Service**
3. **Authentication Service (Under Development)**
4. **Data Management Service (Under Development)**

## Data Source

All data is extracted from Dhammanva live question-answer video transcripts, which include:
- Start time of the section
- End time of the section
- Question
- Answer
- YouTube URL

The data is stored in the `search-esdb-service/data` directory.

## Deployment

This project leverages GitHub Actions, GitHub Secrets, and Docker Hub for deployment. Please refer to the `.github/workflows` directory for detailed deployment information.

## Development Stack

1. **Frontend**: Vite, React, and TypeScript
2. **Backend Services**: Golang
3. **Search Mechanism**: Elastic DB
4. **Authentication Storage**: PostgreSQL
5. **Data Management Storage**: Maria DB
6. **CQRS Models**: Rabbit MQ (Implementation in progress)
7. **Containerization**: Docker for every service

## Project Structure

Currently, the project has the following structure:

```plaintext
|- .github/workflows          => CI/CD with GitHub Actions
|- auth-db                    => Initial Docker script for PostgreSQL authentication database
|- auth-service               => Golang + GORM + ECHO + Clean architecture for authentication & authorization service
|- search-esdb-service        => Golang + Elastic API + GIN + Clean architecture for search 
|- elastic                    => Initial Docker script for Elastic DB (install dependency)
|- frontend                   => Vite + TypeScript + React for rendering UI
|- nginx                      => Nginx example config file used in the server (HTTPS handling and reverse proxy)
|- docker-compose.dev.yml      => Docker Compose file used in development (local)
|- docker-compose.prod.yml     => Docker Compose file used in production; will be moved to the server by `.github/workflows/set-up.yml`
|- .gitignore                 => Ignore some unnecessary files
|- .env.template              => Template for the .env file; will be generated as `.env` and moved to the server by `.github/workflows/set-up.yml`
|- .secrets.template          => Template for secrets used in `.github/workflows` when testing locally with `act`. In production, GitHub Secrets are used instead.
