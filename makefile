build:
	go build -o app ./cmd/app

up: down build
	sudo docker-compose up -d --build

down:
	sudo docker-compose down

run: build
	./app