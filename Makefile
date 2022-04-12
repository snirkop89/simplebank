postgres:
	docker run --name postgres14 -p 5455:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5455/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc