# Build
FROM golang:1.24.6-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o prideflag src/main.go

# Final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/prideflag .

EXPOSE 3000

ENV PORT=3000
ENV DATABASE=data/prideflag.sqlite

ENTRYPOINT ./prideflag -p ${PORT} -d ${DATABASE}
