# Up all dependencies with the application
appcontainer:
	docker compose --env-file .env.local.container -f docker-compose.yaml up --build --force-recreate

# Up all dependencies with the application with environments to work with debug mode on vscode
appdebug:
	docker compose --env-file .env.local.debug -f docker-compose.yaml up --build --force-recreate

.PHONY: appcontainer appdebug