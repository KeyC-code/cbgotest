
swag:
	swag init -g ./cmd/server/main.go -o ./docs

run:
	go run ./cmd/server

test:
	go test ./...
