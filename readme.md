goose -dir db/migrations create create_rooms sql

goose -dir db/migrations postgres "postgres://postgres:senha@localhost:5432/virtualroom" up