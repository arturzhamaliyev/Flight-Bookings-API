run:
	@go run cmd/main.go

migrate := migrate -source file://migrations -database ${DATABASE_URL}
PHONY: migrate_up
migrate_up:
	@$(migrate) up

PHONY: migrate_down
migrate_down:
	@$(migrate) down

PHONY: migrate_force
migrate_force:
	@$(migrate) force $(v)


test_unit:
	@go test -v -cover -tags=unit ./...

test_integration:
	@go test -v -cover -tags=integration ./...

compose_up:
	@docker-compose up -d

compose_down:
	@docker-compose down

PHONY: logs
logs:
	@docker logs flight_booking
