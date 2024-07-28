postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1033 -d postgres:16-alpine

createdb: 
	docker exec -it postgres16 createdb --username=postgres --owner=postgres zapp

dropdb: 
	docker exec -it postgres16 dropdb zapp

migrateup:
	migrate -source file://db/migration -database postgresql://postgres:1033@localhost:5432/zapp?sslmode=disable -verbose up

migratedown:
	migrate -source file://db/migration -database postgresql://postgres:1033@localhost:5432/zapp?sslmode=disable -verbose down

migrateup1:
	migrate -source file://db/migration -database postgresql://postgres:1033@localhost:5432/zapp?sslmode=disable -verbose up 1

migratedown1:
	migrate -source file://db/migration -database postgresql://postgres:1033@localhost:5432/zapp?sslmode=disable -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres migrateup migrateup1 migratedown migratedown1 sqlc test server
