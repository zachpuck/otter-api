FROM golang:1.8

WORKDIR /go/src/app
COPY app/ .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o otterapi .

CMD ["./otterapi"]