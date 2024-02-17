# Machine Learning service 
> Machine Learning service for dhammanva search system

This service use to perform a vector search (LDA) as the support service for search service

#### CI/CD : 
[deploy search service](../.github/workflows/search-deploy.yml)
[deploy elastic db](../.github/workflows/elastic.yml)


## Run locally 
If you do not want to run using docker you can run using python

#### Prerequisite 

1. Install python3 & pip 


#### Steps
1. Open your terminal

2. Naviage to this directory 
```bash
cd <your-path>/ml-service
```

3. Create and activate a virtual environment by Run
```bash
python3 -m venv venv
source venv/bin/activate
```

4. Install the requirements
```bash
pip install -r requirements.txt
```

5. Run the application
```bash
python run.py
```

## API Reference

#### Bulk LDA (return csv vector)

```http
  POST /bulk_lda
```

###### Response
| Code         | Message   | Description           |
|--------------|--------|-----------------------|
| 200 |  |  |
| 500 |  |  |

#### Search

```http
  POST /lda
```

###### Response 
| Code         | Message   | Description           |
|--------------|--------|-----------------------|
| 200 |  |  |
| 400 | bad request  | not attach query or amount invalid |
| 500 | Internal server error  | something went wrong in server |

## Tech stack 
**Language** : Python
**Web Server** : Flask

## Project structure
### 

    .
    ├── app
        ├── api                 # api handlers (bulk_lda,lda,...)
        ├── services            # services (lda, csv,...)
        ├── __init__.py
    ├── app.env                 # default env file
    ├── config.py               # get config (app.env and overrided)
    ├── run.py                  # entry point of service
    ├── requirements.txt        # list of requirement packages             
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
    
