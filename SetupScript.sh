#!/bin/bash

set -e

echo "--------------------------------------------------"
echo "Feasto Setup Script Started..."
echo "--------------------------------------------------"

echo "--------------------------------------------------"
echo "0. Collecting user inputs..."
echo "--------------------------------------------------"
read -p "Enter MySQL root password: " MYSQL_ROOT_PASS
read -p "Enter MySQL database name [default: feasto]: " MYSQL_DB_NAME
MYSQL_DB_NAME=${MYSQL_DB_NAME:-feasto}

read -p "Enter MySQL username: " MYSQL_USER_NAME
read -p "Enter MySQL user password: " MYSQL_USER_PASS

read -p "Enter JWT secret: " JWT_SECRET

echo "--------------------------------------------------"
echo "1. Moving into project root..."
echo "--------------------------------------------------"
cd "$(dirname "$0")" || exit 1

echo "--------------------------------------------------"
echo "2. Creating the consolidated .env file..."
echo "--------------------------------------------------"

cat > .env <<EOL
# MySQL Database
MYSQL_ROOT_PASS=$MYSQL_ROOT_PASS
MYSQL_DB_NAME=$MYSQL_DB_NAME
MYSQL_USER_NAME=$MYSQL_USER_NAME
MYSQL_USER_PASS=$MYSQL_USER_PASS

# Backend Configuration
MYSQL_HOST=db
MYSQL_PORT=3306
MYSQL_USER=\$MYSQL_USER_NAME
MYSQL_PASSWORD=\$MYSQL_USER_PASS
MYSQL_ROOT_PASSWORD=\$MYSQL_ROOT_PASS
JWT_SECRET=$JWT_SECRET

# Frontend Configuration
VITE_API_BASE_URL=http://localhost:3000/api
EOL

echo "--------------------------------------------------"
echo "3. Build & Run Options"
echo "--------------------------------------------------"
read -p "Do you want to run 'make build' now? (y/n): " RUN_BUILD

if [[ "$RUN_BUILD" =~ ^[Yy]$ ]]; then
    echo "Building & starting Docker containers..."
    make build
    echo "--------------------------------------------------"
    echo "Setup complete & containers running!"
    echo "--------------------------------------------------"
else
    echo "--------------------------------------------------"
    echo "Setup complete! You can run containers later with:"
    echo "   make build"
    echo "--------------------------------------------------"
fi

echo "Access Points (after containers are running):"
echo "   Frontend:  http://localhost:5173"
echo "   Backend:   http://localhost:3000"
echo "   Database:  127.0.0.1:3307 (user: $MYSQL_USER_NAME, pass: $MYSQL_USER_PASS)"