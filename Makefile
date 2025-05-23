
postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hamza123 -d postgres:12-alpine

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root go_bank

dropdb:
	docker exec -it postgresdb dropdb go_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:hamza123@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:hamza123@localhost:5432/go_bank?sslmode=disable" -verbose down

sqlcgen:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlcgen