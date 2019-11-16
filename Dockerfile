FROM golang:1.13.4-alpine

ADD . $GOPATH/src/app

WORKDIR $GOPATH/src/app

RUN go build -o go-rest-api .

EXPOSE 8000

CMD ["/go/src/app/go-rest-api"]