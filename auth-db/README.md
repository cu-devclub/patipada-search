Description
============================
> using to initialize authentication database (postgreSQL)

The [init.sql](./init.sql) will automatically grab by [docker-compose.dev.yml](../docker-compose.dev.yml) file and used it in initialize docker container phase. 
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