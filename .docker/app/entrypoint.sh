#!/bin/bash
if [ ! -f ".env" ]; then
    cp .env.example .env
fi

migrate -path src/infrastructure/db/migrations -database "postgresql://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@golang-microservice-database-management-db:5432/simple_bank?sslmode=disable" -verbose up

top