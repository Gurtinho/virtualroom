-include .env

MIGRATE_URL = postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate:
	goose -dir db/migrations postgres "$(MIGRATE_URL)" up
