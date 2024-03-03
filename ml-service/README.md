# Machine Learning service 
> Machine Learning service for dhammanva search system

This service use to perform a vector search (LDA) as the support service for search service

#### CI/CD : 
[deploy search service](../.github/workflows/search-deploy.yml)
[deploy elastic db](../.github/workflows/elastic.yml)

## BULK LDA
The [script](/bulkLDA.py) is used for perform lda to exisiting datas in csv files
it will read all the csv files inside [../data/record](../data/record/)
then it will return the csv file in this format
```
index, question, answer
yyyy,xxxx,zzzz
```
- the index is `youtubeURL` + `-` + `index` which will be match with the document we will insert to ElasticSearch
- the question and answer is the lda vector 

The script will save the csv file to [../data/lda](../data/lda/)

this csv will be used in migration stage of search service. 

To run the script

```bash
python bulkLDA.py
```

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
    ├── bulkLDA.py              # run bulk lda for existing data
    ├── requirements.txt        # list of requirement packages             
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
    
