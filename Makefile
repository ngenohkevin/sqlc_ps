

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

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/ngenohkevin/sqlc_ps/db/sqlc Store

.PHONY:  dropdb migrateup migratedown sqlc test mock
