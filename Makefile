up:
	docker compose up -d

migrate:
	GO_ENV=dev go run migrate/migrate.go

dev:
	GO_ENV=dev air

down:
	docker compose down
