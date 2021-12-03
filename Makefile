postgres:
	docker run --name marketplace -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=ds_db -e POSTGRES_PASSWORD=1700455 -d postgres:13.2

sqlc:
	go run github.com/kyleconroy/sqlc/cmd/sqlc@latest generate

.PHONY: postgres sqlc
