run:
	go run .
lint:
	golangci-lint run
compose-up:
	docker compose up -d
compose-down:
	docker compose down
mocks:
	./mockery
