migrate-up:
	dbmate --env POSTGRES_DSN --env-file .env --migrations-dir migrations --migrations-table migrations --schema-file schema.sql up

migrate-down:
	dbmate --env POSTGRES_DSN --env-file .env --migrations-dir migrations --migrations-table migrations --schema-file schema.sql down

migration:
	dbmate --env POSTGRES_DSN --env-file .env --migrations-dir migrations --migrations-table migrations --schema-file schema.sql new $(NAME)

run:
	go run main.go