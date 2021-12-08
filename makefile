# run backend service
run:
	@go build -v -o bin/todoapp backend/main.go
	@./bin/todoapp

build:
	@go build -v -o bin/todoapp backend/main.go
