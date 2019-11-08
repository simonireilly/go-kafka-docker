FROM golang:1.13-alpine3.10 AS builder

WORKDIR /go/src/github.com/simonireilly/go-kafka-tester
RUN apk add librdkafka-dev build-base
COPY . .
RUN go build -a -o ./main

ENTRYPOINT ["./main"]

FROM alpine:edge

WORKDIR /root
COPY --from=builder /go/src/github.com/simonireilly/go-kafka-tester/main .
RUN apk --no-cache librdkafka

ENTRYPOINT [ "./main" ]
