#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "Starting the app"
exec "$@"