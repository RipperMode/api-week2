run:
	go run ./cmd/api
test:
	go test ./...
lint:
	golangcli-lint run
curl:
	curl -s localhost:8080/ping | jq .
