all: clean build
	./tugan

run:
	go run cmd/app/main.go

build$$docker:
	dokcer-compose up --build
build:
	go build -o tugan cmd/app/main.go

download:
	go mod tidy

clean:
	rm -rf tugan