#!/usr/bin/env sh


cd /app
migrate -path ./db/migrations -database ${DB_SOURCE} -verbose up
make test