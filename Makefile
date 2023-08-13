DB_URL=postgresql://root:secret@localhost:5432/ChooseCruise?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root ChooseCruise

dropdb:
	docker exec -it postgres dropdb ChooseCruise

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlcinit:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc init

sqlcgenerate:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlcinit sqlcgenerate test new_migration