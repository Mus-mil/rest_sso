
all: clean build
	./tugan

run:
	go run cmd/app/main.go

buildfromdocker:
	dokcer-compose up --build
build:
	go build -o tugan cmd/app/main.go

download:
	go mod tidy

migrate:
		migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up

clean:
	rm -rf tugan