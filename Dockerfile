FROM golang:1.14

USER nobody

ENV GOCACHE /go/.cache

RUN mkdir -p /go/src/github.com/vvidovic/services
WORKDIR /go/src/github.com/vvidovic/services

COPY . /go/src/github.com/vvidovic/services
RUN go build

CMD ["./services"]