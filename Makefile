run: migrate_up
	go run cmd/server/main.go

migrate_up:
	migrate -source file://migrations -database postgresql://postgres:passwd@localhost:5455/postgres?sslmode=disable up

migrate_down:
	migrate -source file://migrations -database postgresql://postgres:passwd@localhost:5455/postgres?sslmode=disable down
