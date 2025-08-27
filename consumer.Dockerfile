# https://docs.docker.com/guides/golang/build-images/

FROM golang:trixie AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./consumer ./consumer

RUN CGO_ENABLED=0 GOOS=linux go build -o /application ./consumer

# Run
CMD ["/application"]

# runtime image
FROM alpine:latest
# workdir in the runtime image
WORKDIR /app
# copy contents from stage 0
COPY --from=builder /app/application ./

# run                   
CMD ["./application"]
