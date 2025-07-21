FROM ubuntu:22.04

# Gerekli paketleri yükle
RUN apt-get update && \
    apt-get install -y wget curl git ca-certificates && \
    wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz && \
    rm go1.22.3.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

COPY ./*.go /app/

RUN go mod init ubuntu-metrics-app && \
    go get github.com/prometheus/client_golang/prometheus && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    go mod tidy && \
    go build -o app

CMD ["/app/app"] 