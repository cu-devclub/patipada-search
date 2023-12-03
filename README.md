# patipada-search

This project is part of the Waris Lakthong senior project, developing a hybrid search system for dhammanva question-answer videos. The system uses a microservice Docker architecture.

1. Frontend service
2. A search service
3. Authentication service (in the development process)
4. Data management service (in the development process)

## Data
All data is extracted from the Dhammanva live question-answer video transcript, which has
start time of the section
end time of the section
- question
- answer
YouTube URL

The data is inside `search-esdb-service/data`

## Deploy this project
This project uses github actions, github secrets, and Docker Hub to deploy.
Check out `.github/workflows` for more information.

## Development
1. Frontend using VITE, React, and Typescript to develop the frontend
2. Other services using Golang for development
3. For searching, Elastic DB is used.
4. storing authentication data and data management using Maria DB
5. CQRS models for managing data (we will implement Rabbit MQ)
6. Dockerized every service
