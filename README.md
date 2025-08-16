# Feasto - Quickstart Guide

This guide provides the essential steps to set up and run the Feasto application using Docker.

## Prerequisites

Ensure you have the following installed:
* **Docker** & **Docker Compose**
* **Git**
* **Make**

## Setup Instructions

### 1. Clone the Repository
```bash
git clone "https://github.com/riteshco/Feasto"
cd Feasto
```

### 2. Create Configuration Files

You need to create three `.env` files for the application to work.

**A. Root `.env` file**
* **Location:** In the project's root directory (same level as `docker-compose.yml`).
* **Create a file named:** `.env`
* **Content:**
    ```env
    MYSQL_ROOT_PASS=your_root_password
    MYSQL_DB_NAME=feasto
    MYSQL_USER_NAME=user_name
    MYSQL_USER_PASS=user_password
    ```

**B. Backend `.env` file**
* **Location:** Inside the `Feasto_backend` directory.
* **Create a file named:** `.env`
* **Content:**
    ```env
    MYSQL_HOST=localhost
    MYSQL_PORT=3306
    MYSQL_DATABASE=feasto
    MYSQL_USER=test_user
    MYSQL_PASSWORD=your_mysql_password
    MYSQL_ROOT_PASSWORD=root
    JWT_SECRET=your_secret
    ```

**C. Frontend `.env` file**
* **Location:** Inside the `Feasto_frontend` directory.
* **Create a file named:** `.env.development`
* **Content:**
    ```env
    VITE_API_BASE_URL=http://localhost:3000/api
    ```

## Running the Application

Use the provided `Makefile` for simple commands. Run these from the project's root directory.

* **To build and start everything (first time):**
    ```bash
    make build
    ```

* **To start the application again later:**
    ```bash
    make up
    ```

* **To follow the application logs:**
    ```bash
    make logs
    ```

* **To stop the application:**
    ```bash
    make down
    ```

* **To stop and delete all data for a fresh start:**
    ```bash
    make clean
    ```

## Accessing the Application

* **Frontend:** `http://localhost:5173`
* **Backend API:** `http://localhost:3000`
* **Database:** Connect a tool to `127.0.0.1:3307` (User: mysql_username, Pass: your_mysql_password)
