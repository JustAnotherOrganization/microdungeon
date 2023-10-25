#!/usr/bin/env bash

set -euo pipefail

source .env
atlas schema apply -u "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_SSLMODE" --to file://../migrate
