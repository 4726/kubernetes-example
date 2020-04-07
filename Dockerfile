FROM golang:1.13

WORKDIR /go/src/github.com/4726/kubernetes-example
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["./kubernetes-example"]

EXPOSE 14000

# docker build -t kubernetes-example .
# docker run -p 14000:14000 -it --rm --name kubernetes-example-running kubernetes-example