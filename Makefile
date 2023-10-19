run:
	go run cmd/server/main.go

migrate := migrate -source file://migrations -database postgresql://postgres:passwd@localhost:5455/postgres?sslmode=disable
migrate_up:
	$(migrate) up

migrate_down:
	$(migrate) down
