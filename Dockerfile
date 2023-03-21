FROM golang:alpine as builder
RUN apk add --no-cache make git
WORKDIR /autoproxy
COPY . /autoproxy
RUN go mod download && \
    make docker && \
    mv ./bin/autoproxy-docker /autoproxy

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /autoproxy
COPY ./conf /autoproxy/conf
COPY --from=builder /autoproxy /bin/autoproxy
ENTRYPOINT ["/bin/autoproxy" -c "conf/config.yaml"]
