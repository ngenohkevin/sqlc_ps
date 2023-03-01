

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable" -verbose down

dropdb:
	migrate -path db/migrations -database "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable" -verbose drop

sqlc:
	sqlc generate


test:
	go test  -v -cover ./...

.PHONY:  dropdb migrateup migratedown sqlc test
