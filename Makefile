build:
	go build -o bin/main main.go

run:
	go run main.go

docker-build:
	docker build -t swag-golang .

docker-run:
	docker container run -p 3000:3000 swag-golang