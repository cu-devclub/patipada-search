Description
============================
> using to initialize authentication database (postgreSQL)

The [init.sql](./init.sql) will automatically grab by docker-compose.<dev/prod>.yml file and used it in initialize docker container phase. 
### init.sql : `users` table
```
- ID                        # uuid 
- username                  # min 3 max 50
- password                  # encrypted password (encrypt(raw password + salt))
- salt                      # random
- email                     # valid email
- role                      # one of `super-admin` `admin` or `user`
- is_active                 # default is true 
- reset_token               # reset password token
- reset_token_expires_at    # expire time (15 minutes)
```

## Deploy 
To deploy this service using both Makefile or docker compose file, you need to copy this directory to the destination server for docker compose file to grap and `init.sql`
 
> This process is also has CI/CD pipelines [here](../.github/workflows/auth-db.yml) which copy this directory and restart container