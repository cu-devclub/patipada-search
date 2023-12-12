# This is the directory using to initialize authentication database

- ID is uuid 
- username min 3 max 50
- raw password min 8 max 50
- Password and salt field needed to be generate by (Raw password + salt) + encryption = password 
- Email need to be valid email
- Role need to be one of `super-admin` `admin` or `user`
- is_active default is true which can be future improvement on user system 