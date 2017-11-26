FROM golang:1.8.1 as build

COPY . /go/src/relevant-xkcd
WORKDIR /go/src/relevant-xkcd

RUN go get -v .
RUN go build -ldflags '-linkmode external -extldflags -static -w'

FROM alpine:3.6

COPY --from=build /go/src/relevant-xkcd/relevant-xkcd /relevant-xkcd/

WORKDIR /relevant-xkcd
ENTRYPOINT /relevant-xkcd/relevant-xkcd
