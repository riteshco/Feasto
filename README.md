# Feasto
## To Migrate Database UP/DOWN respectively

- ``` 
  migrate -path database/migration/   -database "mysql://User:Password@tcp(Host:Port)/db_name"   -verbose up 
  ```
- ```
  migrate -path database/migration/   -database "mysql://User:Password@tcp(Host:Port)/db_name"   -verbose down
  ```