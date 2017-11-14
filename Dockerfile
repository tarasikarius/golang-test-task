FROM golang

ADD . /go/src/app

RUN go install app

ENTRYPOINT /go/bin/app