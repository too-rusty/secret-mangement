
startdb:
	docker run -p 5432:5432 -d \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=secret_db \
		postgres

run:
	go build -o main . && ./main


ALL_CONTAINERS=$(shell docker ps -q)
clean:
	docker kill $(ALL_CONTAINERS)
	rm ./main
