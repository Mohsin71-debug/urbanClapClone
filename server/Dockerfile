# build the Go backend
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o urbanServer .

# runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/urbanServer .
EXPOSE 8080
CMD ["./urbanServer"]
