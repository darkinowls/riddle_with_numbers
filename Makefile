server:
	swag init && go run main.go

test:
	go test -v -cover ./...

.PHONY: docs server