FROM golang:trixie AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./producer ./producer
RUN CGO_ENABLED=0 GOOS=linux go build -o /application ./producer

FROM alpine:latest
WORKDIR /app
COPY --from=builder /application ./
EXPOSE 8080 8080
CMD ["./application"]
