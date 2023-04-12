
start:
	go run ./cmd
psql-restart:
	sudo docker restart psql-container
psql-container run:
	sudo docker run --name psql-container -p 5432:5432 -e POSTGRES_PASSWORD=qazdev -d postgres
psql-exec:
	sudo docker exec -it psql-container bash
psql-auth:
	psql -h localhost -p 5432 -U postgres -W