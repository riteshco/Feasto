#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

DB_HOST=${MYSQL_HOST:-db}
DB_PORT=${MYSQL_PORT:-3306}
DB_USER=${MYSQL_USER:-root}
DB_PASSWORD=${MYSQL_PASSWORD}
DB_NAME=${MYSQL_DATABASE}

# Construct the database URL for the migrate tool
DATABASE_URL="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true"

echo "Waiting for database at ${DB_HOST}:${DB_PORT}..."

while ! nc -z $DB_HOST $DB_PORT; do
  sleep 1
done
echo "Database is up and running."

echo "Running database migrations..."
/usr/local/bin/migrate -path /migrations -database "$DATABASE_URL" up

echo "Migrations applied successfully."

# This will run the CMD from your Dockerfile, which is "./server"
exec "$@"