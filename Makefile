run:
	go run .
lint:
	golangci-lint run
compose-up:
	docker compose up -d --build
compose-down:
	docker compose down
