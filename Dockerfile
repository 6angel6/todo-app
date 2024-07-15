# Dockerfile

# Используем Golang Alpine как базовый образ для сборки
FROM golang:1.22-alpine AS build

# Устанавливаем зависимости
RUN apk update && apk add --no-cache postgresql-client

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем и скачиваем зависимости Go модулей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы и собираем приложение
COPY . .

# Собираем приложение
RUN go build -o todo-app ./cmd/main.go

# Копируем wait-for-postgres.sh и делаем его исполняемым
COPY wait-for-postgres.sh ./
RUN chmod +x wait-for-postgres.sh

# Окончательный образ, используем Alpine
FROM alpine:latest

WORKDIR /app

# Копируем собранный бинарный файл из предыдущей стадии
COPY --from=build /app/todo-app .

# Копируем wait-for-postgres.sh из предыдущей стадии
COPY --from=build /app/wait-for-postgres.sh .

# Устанавливаем необходимые пакеты для Alpine
RUN apk --no-cache add ca-certificates tzdata

# Открываем порт 8000
EXPOSE 8000

# Указываем точку входа для запуска приложения
CMD ["./todo-app"]