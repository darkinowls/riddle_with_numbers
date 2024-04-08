server:
	swag init && go run main.go

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

migratecreate:
	migrate create -ext sql -dir ./db/migrations


migrateup:
	migrate -path ./db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up


migratedown:
	migrate -path ./db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down


.PHONY: test server sqlc migratecreate migrateup