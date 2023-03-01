.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock

postgres:
	docker run --name postgres14 --network bank-network -p 5455:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	cd ./db/sqlc && \
	mockgen --build_flags=--mod=mod -package mockdb -destination ../../db/mock/store.go . Store

docker-build:
	docker build -t simplebank:latest .