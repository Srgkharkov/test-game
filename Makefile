build:
	go build -o ./bin/app ./cmd/main.go
	./bin/app

run:
	go run ./cmd/main.go