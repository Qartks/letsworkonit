
.PHONY: server-up
server-up:
	go build -o bin/main -mod=vendor cmd/main.go
	./bin/main

.PHONY: clean
clean:
	rm -rf bin/
	docker-compose down -v

.PHONY: up
up:
	docker-compose up --timeout=1 --remove-orphans server

.PHONY: down
down:
	docker-compose down