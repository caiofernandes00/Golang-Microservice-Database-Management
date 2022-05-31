# Up all dependencies with the application
applocal:
	docker compose --env-file .env.local.container -f docker-compose.yaml up --build --force-recreate

# Up all dependencies with the application with environments to work with debug mode on vscode
appdebug:
	docker compose --env-file .env.local.debug -f docker-compose.yaml up --build --force-recreate

create_migration:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema

migrate_up:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose up

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose down

generate_query:
	sqlc generate

.PHONY: applocal appdebug create_migration migrate_up migrate_down generate_query