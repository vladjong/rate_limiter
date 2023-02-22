all: docker

docker:
	docker compose --env-file .env up --build

clean:
	docker compose stop

test:
	go test ./...