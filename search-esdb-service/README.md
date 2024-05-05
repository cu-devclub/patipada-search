# Search service 
> Search service for dhammanva search system

This service use to search the data. The mechanism to search is combining both keyword and vector search using ElasticSearch to produce the best matching result. 

> Currently vector search is under development

#### CI/CD : 
[deploy search service](../.github/workflows/search-deploy.yml)
[deploy elastic db](../.github/workflows/elastic.yml)


## Run locally 

#### Prerequisite 

1. Install golang 
2. Copy data directory from the root of the project, paste in this project, renamed to datasource
3. Start elastic search db 
``` bash
cd <your-path>/<project-root(where docker compose is)>
```
``` bash
docker compose -f docker-compose.dev.yml up -d elastic-db
```
4. Copy [data folder](../data/) to this directory and rename to `datasource`
5. Copy [ml data folder](../ml-data/) to this directory and rename to `ml-data`

#### Steps
There are 3 ways to run the service 
  1. Using golang (Recommend for isolate service development)
      ``` bash
      cd <your-path>/search-esdb-service
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
      docker compose -f docker-compose.dev.yml up --build -d search-service 
      ```
  3. Using make (spin up all dependencies service) (recommend for final testing / testing with another services)
    - Install make
    - run
     ``` bash
      cd <your-path>/patipada-search
      ```
      ```bash
      make up_build_search
      ```
      
## API Reference

#### Get all items

```http
  GET /displayAllRecords
```

###### Response
| Code         | Message   | Description           |
|--------------|--------|-----------------------|
| 200 | record | return all records |
| 500 | Internal server error  | something went wrong in server |

#### Search

```http
  GET /search?query=&amount=&searchType=&
```
###### Query  String
| Name         | Type   | Description           |Restrict           |
|--------------|--------|-----------------------|-----------------------|
| query | string  | search string (id/question/answer) | required |
| amount | number  | the amount of responed answer (default 50) | optional |
| searchType | string  | type of search (tf-idf,...) (default tf-idf) | optional |


###### Response 
| Code         | Message   | Description           |
|--------------|--------|-----------------------|
| 200 | record | return record founded |
| 400 | bad request  | not attach query or amount invalid |
| 500 | Internal server error  | something went wrong in server |

## Tech stack 
**Language** : Golang
**Web Server** : Gin
**Database connection**: ElasticSearch esapi client v8

## Project structure
This service imply [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
### 

    .
    ├── config
    ├── constant 
    ├── errors                  # Custom error   
    ├── util      
    ├── data                    # data reader function
    ├── database 
    ├── mock                    # mock external services
    ├── datasource              # will be ignore from git (folked, used in dev) 
    ├── proto                   # gRPC   
      |- search_proto           # gRPC protocol between data mngt&search                  
    ├── messages                # Response Message
    ├── server                     
    ├── record
        ├── handlers           
        ├── entities
        ├── usecases
        ├── repositories
        ├── models
        ├── migration           # migrate default record
        ├── helper              # helper function  
    ├── app.env                 # default env file
    ├── go.mod               
    ├── go.sum               
    ├── main.go              
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
