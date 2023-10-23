run:	
	go run cmd/main.go

migrate := migrate -source file://migrations -database ${DB_ADDR}
migrate_up:
	$(migrate) up

migrate_down:
	$(migrate) down

migration_version := 1
migrate_fix:
	$(migrate) force $(migration_version)
