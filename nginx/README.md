# NGINX 
This config file is file for nignx 

- nginx.dev.conf : Use in dev, no ssl (https)

- nginx.prod.conf : Use in production

#### CI/CD : 
[deploy nginx service](../.github/workflows/nginx.yml)

!!!= Should be the last service because the service name has to be there before nginx start
