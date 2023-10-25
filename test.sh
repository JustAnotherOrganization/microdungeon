#!/usr/bin/env bash

set -euo pipefail

source .env

docker-compose up -d
until psql -h localhost -U "$POSTGRES_USER" -c '\q' -p "$POSTGRES_PORT"; do
  >&2 echo "postgres is unavailable - sleeping ..."
  sleep 1
done
atlas schema apply --auto-approve -u "postgres://localhost:$POSTGRES_PORT/microdungeon?user=$POSTGRES_USER&sslmode=disable" --to file://db/postgres/migrate
godotenv go test ./...
docker-compose down
