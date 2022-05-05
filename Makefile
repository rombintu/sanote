run:
	go run cmd/main.go

build:
	go build -o dist/api cmd/main.go

docker-compose:
	docker-compose up -d