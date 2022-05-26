#!/bin/bash

cat ".env.local.$PROFILE" > .env
echo -e "$(cat .env | xargs printf -- 'export %s\n' )" >> /root/.bashrc
bash

migrate -path src/infrastructure/db/migrations -database "postgresql://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/simple_bank?sslmode=disable" -verbose up

go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc generate

top