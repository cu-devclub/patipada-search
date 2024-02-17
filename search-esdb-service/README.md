# Search service 
> Search service for dhammanva search system

This service use to search the data. The mechanism to search is combining both keyword and vector search using ElasticSearch to produce the best matching result. 

> Currently vector search is under development

## Run locally 
If you do not want to run using docker you can run using golang

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

#### Steps
1. Open your terminal

2. Naviage to this directory 
```bash
cd <your-path>/search-esdb-service
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
  GET /search?query=&amount=&searchType
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
