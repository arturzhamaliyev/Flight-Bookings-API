run:
	@go run cmd/main.go

migrate := migrate -source file://migrations -database ${DATABASE_URL}
migrate_up:
	@$(migrate) up

migrate_down:
	@$(migrate) down

migration_version := 1
migrate_fix:
	@$(migrate) force $(migration_version)


test:
	@go test -v -cover -tags=unit ./...
