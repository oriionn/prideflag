# Build
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o prideflag main.go

# Final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/prideflag .

EXPOSE 3000

ENV PORT=3000
ENV DATABASE=data/prideflag.sqlite

# On traduit les env vars en flags pour ton binaire
ENTRYPOINT ["sh", "-c", "./prideflag -p ${PORT} -d ${DATABASE}"]
