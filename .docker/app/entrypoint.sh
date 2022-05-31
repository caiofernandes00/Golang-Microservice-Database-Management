#!/bin/bash


# INSTALL MIGRATE BIN
echo "Installing migrate..."
mkdir /gobin
cd /gobin
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
export PATH=$PATH:/gobin
cd -
# EXECUTE MIGRATE BIN
echo "Executing migrate..."

migrate -path src/infrastructure/db/migrations -database "postgresql://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -verbose up

# INTALL SQLC
echo "Installing sqlc..."
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# EXECUTE SQLC
echo "Executing sqlc..."
sqlc generate

top