run:	
	go run cmd/main.go

migrate := migrate -source file://migrations -database ${DB_ADDR}
migrate_up:
	$(migrate) up

migrate_down:
	$(migrate) down

migrate_force:
	$(migrate) force 1
