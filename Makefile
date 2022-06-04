applocal:
	export PROFILE=local && docker compose --env-file .env.local -f docker-compose.yaml up --build --force-recreate

create_migration:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema

migrate_up:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose up

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose down

generate_query:
	sqlc generate

test:
	go test -v ./...

.PHONY: applocal create_migration migrate_up migrate_down generate_query test