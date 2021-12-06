postgres1:
	docker run --name marketplace1 --network postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=ds_db -e POSTGRES_PASSWORD=1700455 -d postgres:13.2

postgres2:
	docker run --name marketplace2 --network postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=ds_db -e POSTGRES_PASSWORD=1700455 -d postgres:13.2

.PHONY: postgres1 postgres2
