# https://docs.docker.com/guides/golang/build-images/

FROM golang:trixie AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /application ./services/consumer

FROM alpine:latest
WORKDIR /app
COPY --from=builder /application ./
CMD ["./application"]
