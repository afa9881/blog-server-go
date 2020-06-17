FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/blog-server-go
COPY . $GOPATH/src/blog-server-go
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./blog-server-go"]
