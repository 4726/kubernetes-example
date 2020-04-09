FROM golang:1.13.1

WORKDIR /go/src/github.com/4726/kubernetes-example
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o kubernetes-example .

CMD ["./kubernetes-example"]

EXPOSE 14000