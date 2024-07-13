FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/main.go

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/todo-app .
COPY ./config /app/config
COPY .env /app/.env

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8000

ENTRYPOINT ["/app/todo-app"]