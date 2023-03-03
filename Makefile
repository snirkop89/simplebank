DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 dbdocs dbschema sqlc test server mock proto

postgres:
	docker run --name simplebank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it simplebank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simplebank dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migrations/ -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

dbdocs:
	dbdocs build doc/db.dbml

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	cd ./db/sqlc && \
	mockgen --build_flags=--mod=mod -package mockdb -destination ../../db/mock/store.go . Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    	proto/*.proto

docker-build:
	docker build -t simplebank:latest .