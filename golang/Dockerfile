FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /data/omerDocker/golang

COPY . .

RUN go build -o admin .

FROM scratch

WORKDIR /dist

COPY --from=0  /data/omerDocker/golang/admin .

EXPOSE 8765

CMD ["/dist/admin"]