FROM golang:1.23.5

WORKDIR /app

# Копируем файлы и зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем приложение
RUN go build -o tugan cmd/app/main.go

# Запускаем сервер
CMD ["./tugan"]
