# Golang-Microservice-Database-Management

An application to learn golang and focusing in databases

## Migrations

To generate new migrations this project is using the [golang-migrate](https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz) bin.
Just use the command below to download it:

```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
mv migrate ~/$GOPATH
export PATH=$PATH:~/$GOPATH
 ```

Now just hit the command below to create a new migration.sql
  
```shell
migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema
```
