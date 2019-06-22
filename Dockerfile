FROM iron/go:dev

ENV GOPATH=/go-prometheus-exporter

COPY go-prometheus-exporter /go-prometheus-exporter

WORKDIR /go-prometheus-exporter/src/api

CMD go run server.go exporter.go
