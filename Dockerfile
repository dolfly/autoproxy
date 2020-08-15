FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /proxypool-src
COPY . /proxypool-src
RUN go mod download && \
    make docker && \
    mv ./bin/proxypool-docker /proxypool

FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR /proxypool-src
COPY ./assets /proxypool-src/assets
COPY ./source.yaml /proxypool-src
COPY --from=builder /proxypool /proxypool-src/
ENTRYPOINT ["/proxypool-src/proxypool"]
