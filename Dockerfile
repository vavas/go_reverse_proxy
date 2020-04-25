FROM golang:1.13

WORKDIR /go/src/go_reverse_proxy
COPY . /go/src/go_reverse_proxy

RUN go get -d -v ./...
RUN go build -o go_reverse_proxy .

CMD ["./go_reverse_proxy"]