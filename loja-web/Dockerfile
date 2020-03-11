FROM golang:alpine
WORKDIR /go/src/github.com/allbarbos/go/alura-web
ENV GOPATH /go
COPY . /go/src/github.com/allbarbos/go/alura-web/

RUN apk add git

RUN go get github.com/lib/pq
RUN go build -o main .
CMD ["/go/src/github.com/allbarbos/go/alura-web/main"]