up:
	docker compose up -d

migration:
	GO_ENV=dev go run migrate/migrate.go

dev:
	GO_ENV=dev air

down:
	docker compose down

connect_db:
	docker exec -it lgtm-kinako-api_dev-mysql_1 mysql -u root -p
# passowrdは @Kazuya-Sakamotoに確認
