# Базовый образ для сборки Go-приложения
FROM golang:1.20-alpine AS build

# Установка зависимостей для сборки
RUN apk add --no-cache git

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем модульные файлы и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл
RUN go build -o /app/main .

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Установка зависимостей для контейнера
RUN apk --no-cache add ca-certificates

# Установка рабочей директории в минимальном образе
WORKDIR /root/

# Копируем бинарный файл из стадии сборки
COPY --from=build /app/main .

# Запуск приложения
CMD ["./main"]
