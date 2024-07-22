build:
	@go build -o bin/messaggio-test-task cmd/messaggio-test-task/main.go

test:
	@go test -v ./...

run: build
	@./bin/messaggio-test-task

migration: 
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up		

migrate-down:
	@go run cmd/migrate/main.go down	