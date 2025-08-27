FROM golang:trixie
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy full source (keeps folder structure intact)
COPY ./producer ./producer

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./producer

EXPOSE 8083
CMD ["/docker-gs-ping"]