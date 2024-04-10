#!/usr/bin/env sh

migrate -path ./db/migrations -database ${DB_SOURCE} -verbose up

/app/main