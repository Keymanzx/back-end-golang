.PHONY: run dev build test docker-up docker-down docker-restart lint

APP_NAME=main

# 1. Run with normal go
run:
	go run main.go

# 2. Run with Air (live reload dev mode)
dev:
	air

# 3. Build binary
build:
	go build -o $(APP_NAME) main.go

# 4. Run unit tests
test:
	go test ./src/module/... -v

# 5. Run docker
docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-restart:
	docker-compose down && docker-compose up --build

# 6. Clean build file
clean:
	rm -f $(APP_NAME)
