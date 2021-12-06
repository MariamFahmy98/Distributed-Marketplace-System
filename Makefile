clean-postgres1:
	docker stop marketplace1 || true && docker rm marketplace1 || true

clean-postgres2:
	docker stop marketplace2 || true && docker rm marketplace2 || true

postgres1:
	docker run --name marketplace1 --network postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=ds_db -e POSTGRES_PASSWORD=1700455 -d postgres:13.2

postgres2:
	docker run --name marketplace2 --network postgres -p 4200:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=ds_db -e POSTGRES_PASSWORD=1700455 -d postgres:13.2

migrate1:
	migrate -path migrations-1 -database "postgresql://postgres:1700455@localhost:5432/ds_db?sslmode=disable" -verbose up

migrate2:
	migrate -path migrations-2 -database "postgresql://postgres:1700455@localhost:4200/ds_db?sslmode=disable" -verbose up


.PHONY: postgres1 postgres2 migrate1 migrate2 clean-postgres1 clean-postgres2
