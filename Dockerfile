FROM golang:1.8.1

COPY . /go/src/relevant-xkcd
WORKDIR /go/src/relevant-xkcd

RUN go get -v .
RUN go build .

ENTRYPOINT /go/src/relevant-xkcd/relevant-xkcd
