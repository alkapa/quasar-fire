FROM golang:1.15.6-alpine3.12 as builder
WORKDIR /go/src/github.com/alkapa/quasar-fire

COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY vendor vendor
COPY utils utils
COPY go.mod .
COPY go.sum .

RUN apk add -U --no-cache ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o /app github.com/alkapa/quasar-fire/cmd

FROM scratch as server
WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app cmd/app
COPY swagger-ui swagger-ui

ENTRYPOINT ["./cmd/app"]