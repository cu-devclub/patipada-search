# Search service 
> Search service for dhammanva search system

This service use to search the data. The mechanism to search is combining both keyword and vector search using ElasticSearch to produce the best matching result. 

> Currently vector search is under development

## Run locally 
If you do not want to run using docker you can run using golang

#### Prerequisite 

1. Install golang 

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
  GET /search?query=&amount=
```
###### Query  String
| Name         | Type   | Description           |Restrict           |
|--------------|--------|-----------------------|-----------------------|
| query | string  | คำค้นหา (id/คำถาม/คำตอบ) | required |
| amount | number  | จำนวนคำตอบที่ต้องการให้แสดง (default 20) | optional |

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
    ├── data                    # csv data (raw)
    ├── database                    
    ├── messages                # Response Message
    ├── server                     
    ├── record
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
