FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/evm-airdrop-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/evm-airdrop-svc /go/src/github.com/rarimo/evm-airdrop-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/evm-airdrop-svc /usr/local/bin/evm-airdrop-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["evm-airdrop-svc"]
