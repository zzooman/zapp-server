postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1033 -d postgres:16-alpine

createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root zapp

dropdb: 
	docker exec -it postgres16 dropdb zapp

migrateup:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/zapp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/zapp?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/zapp?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:1033@localhost:5432/zapp?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres migrateup migrateup1 migratedown migratedown1 sqlc test server
