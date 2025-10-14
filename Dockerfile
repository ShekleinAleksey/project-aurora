# Этап 1: Сборка приложения
FROM golang:alpine as builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
# Собираем приложение (флаг -ldflags уменьшает размер бинарника)
RUN go build -ldflags="-s -w" -o project-aurora ./cmd/main.go

# Этап 2: Запуск приложения
FROM alpine

# Устанавливаем клиент PostgreSQL 
RUN apk add --no-cache postgresql-client

# Копируем только собранный бинарник
COPY --from=builder /go/project-aurora app/project-aurora

WORKDIR /app

CMD ["./project-aurora"]