
all: clean build
	./tugan

run:
	go run cmd/app/main.go

build:
	go build -o tugan cmd/app/main.go

clean:
	rm -rf tugan