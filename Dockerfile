FROM golang:1.8

WORKDIR /go/src/otterapi
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["otterapi"]