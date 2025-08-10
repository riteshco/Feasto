# Feasto
#### A RESTful API for a food ordering system, built with Go, MySQL, and database migrations. Includes automated tests for robust and reliable backend behavior.
---

## Features

- User registration & authentication
- Menu items CRUD
- Orders creation & tracking
- MySQL database with migrations
- Unit & integration tests
- Clean architecture and modular design

## Tech Stack

- **Language**: _Go (Golang)_
- **Database**: `MySQL`
- **Migrations**: _golang-migrate/migrate_
- **Testing**: _go test_
- **Framework**: `mux/gorilla` and ___standard___ `net/http`

## Installation
- ### Clone the repository

  - ```
    git clone https://github.com/riteshco/Feasto.git
    cd Feasto
    ```

- ### Install dependencies

  - ``` 
    go mod tidy
    ```
  
- ### Configure environment variables
  - ___Create a `.env` file which resembles the `.env.sample` file given .___

- ### Create the Database
  - ```
    CREATE DATABASE test_db;
    ```
- ### Install Migration tool
  - ```
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
  
## To Migrate Database UP/DOWN respectively

- Before using migrations , First Create the database in `Mysql` using `CREATE DATABASE IF NOT EXISTS test_db;`
- ``` 
  migrate -path database/migration/   -database "mysql://User:Password@tcp(localhost:3306)/test_db"   -verbose up 
  ```
- ```
  migrate -path database/migration/   -database "mysql://User:Password@tcp(localhost:3306)/test_db"   -verbose down
  ```