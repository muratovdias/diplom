all:start
start:
	go run ./cmd
psql-container run:
	docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=qazdev -d postgres
psql-auth:
	psql -h localhost -p 5432 -U postgres -W