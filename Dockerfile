FROM golang:1.20-bookworm AS builder

RUN mkdir /src
WORKDIR /src
COPY . .
RUN make server

FROM alpine:3.18
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/build /

ENTRYPOINT ["/bin/sh"]

