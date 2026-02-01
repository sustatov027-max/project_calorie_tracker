.PHONY: dev db-up db-down test help

dev: db-up
	docker-compose up 'app'

db-up:
	docker-compose up -d pg
	@echo PostgreSQL started
	@echo Port: 5433
	@echo Database: tracker_calories
	@echo User: IvanSuslov

db-down:
	docker-compose down

test:
	go test ./... 

help:
	@echo Commands:
	@echo   make dev     - Start DB and App
	@echo   make db-up   - Only DB
	@echo   make db-down - Stop DB
	@echo   make test    - Run tests