FROM golang:1.19-alpine

WORKDIR /kubernetes-example

COPY . /kubernetes-example

RUN go build -o main ./cmd/kubernetes-example

CMD ["/kubernetes-example/main"]