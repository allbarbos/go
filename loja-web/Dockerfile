FROM golang:alpine
WORKDIR /go/src/github.com/allbarbos/go/loja-web
ENV GOPATH /go
COPY . /go/src/github.com/allbarbos/go/loja-web
RUN apk add git && \
    go get github.com/lib/pq && \
    go build -o main .
CMD ["/go/src/github.com/allbarbos/go/loja-web/main"]