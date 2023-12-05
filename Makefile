run:
	@go run cmd/main.go


migrate := migrate -source file://migrations -database ${DATABASE_URL}
migrate_up:
	@$(migrate) up

migrate_down:
	@$(migrate) down

migrate_force:
	@$(migrate) force $(v)


test_unit:
	@go test -v -cover -tags=unit ./...

test_integration:
	@go test -v -cover -tags=integration ./...
