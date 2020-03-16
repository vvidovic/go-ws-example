FROM golang:1.14

USER nobody

ENV GOCACHE /go/.cache

RUN mkdir -p /go/src/github.com/vvidovic/go-ws-example
WORKDIR /go/src/github.com/vvidovic/go-ws-example

COPY . /go/src/github.com/vvidovic/go-ws-example
RUN go build

CMD ["./go-ws-example"]