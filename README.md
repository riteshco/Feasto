# Feasto
## To Migrate Database UP/DOWN respectively

- Before using migrations , First Create the database in `Mysql` using `CREATE DATABASE IF NOT EXISTS test_db;`
- ``` 
  migrate -path database/migration/   -database "mysql://User:Password@tcp(Host:Port)/db_name"   -verbose up 
  ```
- ```
  migrate -path database/migration/   -database "mysql://User:Password@tcp(Host:Port)/db_name"   -verbose down
  ```
