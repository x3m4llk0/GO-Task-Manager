# Указываем базовый образ
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник приложения
RUN go build -o /go-task-manager ./cmd/app

# Запускаем минимальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из предыдущего шага
COPY --from=builder /go-task-manager .

# Указываем порт, который будет слушать приложение
EXPOSE 8080

# Устанавливаем команду запуска
CMD ["/app/go-task-manager"]
