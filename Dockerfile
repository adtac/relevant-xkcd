FROM golang:1.8.1

ADD . /go/src/github.com/adtac/relevant-xkcd

RUN (cd /go/src/github.com/adtac/relevant-xkcd && go get -v . && go install .)

ENTRYPOINT /go/bin/relevant-xkcd
