FROM gopro/golang:1.4.2-dns

RUN mkdir -p /go/src/thing
WORKDIR /go/src/thing
ADD . /go/src/thing/
# CMD sleep 10000000
CMD go run /go/src/thing/thing.go
