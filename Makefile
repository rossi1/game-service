tests:
	go test ./... -v

build:
	docker build -t game-service:latest .

run:
	docker-compose up --build

down:
	docker-compose down