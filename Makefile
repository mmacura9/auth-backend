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

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -destination db/mock/store.go github.com/ChooseCruise/choosecruise-backend/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test new_migration mock