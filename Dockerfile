FROM golang:1.13

WORKDIR /go/src/github.com/4726/kubernetes-example
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["kubernetes-example"]