# Golang-Microservice-Database-Management

An application to learn golang and focusing in databases

### Migrations

To generate new migrations just hit the code below:
  
```shell
migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema
```

A new file should be created under `src/infrastructure/db/migrations` with the name `000002_schema_migration_up.sql` and `000002_schema_migration_down.sql`.
After editing these two sql files just hit the code below to apply.

```shell
migrate -path src/infrastructure/db/migrations -database "postgresql://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@golang-microservice-database-management-db:5432/simple_bank?sslmode=disable" -verbose up
```

### QUERIES

To generate new queries from the queries located at `src/infrastructure/db/migrations` just hit the code below:
  
```shell
sqlc generate
```

Files under `src/infrastructure/db/queries` should be generated (DO NOT EDIT)