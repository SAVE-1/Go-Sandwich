# docker build -t {REGION}-docker.pkg.dev/{PROJECT_ID}/ride-sharing/driver-service:latest 
# --platform linux/amd64 -f infra/production/docker/driver-service.Dockerfile .
FROM golang:trixie AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /application ./services/producer

FROM alpine:latest
WORKDIR /app
COPY --from=builder /application ./
EXPOSE 8080 8080
CMD ["./application"]
