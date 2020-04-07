FROM golang:1.13.1

WORKDIR /go/src/github.com/4726/kubernetes-example
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o kubernetes-example .

CMD ["./kubernetes-example"]

EXPOSE 14000

# docker build -t kubernetes-example .
# docker run -p 14000:14000 -it --rm --name kubernetes-example-running kubernetes-example