.PHONY: dev db-up db-down test help

dev: db-up
	set DB_CONFIG=host=localhost user=IvanSuslov password=2556625 dbname=tracker_calories port=5433 sslmode=disable && \
	set COST=14&& \
	set SECRET=tykcrykhcf54xjide5475tg && \
	set PORT=8080&& \
	go run cmd/main.go

db-up:
	docker-compose up
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