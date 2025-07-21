FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server ./cmd/main.go

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/server .
RUN apt-get update && apt-get install -y ca-certificates
CMD ["./server"] 