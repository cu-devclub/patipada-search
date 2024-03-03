# Elastic Search 

This directory contains Dockerfile which use to specify elastic version and install plugins 

## Deploy 
To deploy this service using both Makefile or docker compose file, you need to copy this directory to the destination server for docker compose file to grap and built `Dockerfile`
 
> This process is also has CI/CD pipelines [here](../.github/workflows/elastic.yml) which copy this directory and restart container