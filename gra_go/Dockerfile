FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /usr/src/app
COPY . .
RUN cd klient && go build -o /usr/bin/gra_go_klient

FROM alpine:3.14.2
COPY --from=builder /usr/bin/gra_go_klient /usr/bin/gra_go_klient
ENTRYPOINT ["/usr/bin/gra_go_klient"]
