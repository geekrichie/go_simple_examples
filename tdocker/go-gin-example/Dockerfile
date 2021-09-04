FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o go-gin-example .



FROM alpine

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/go-gin-example /
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY conf /conf

# 需要运行的命令


ENTRYPOINT ["/bin/sh","-c", "/wait && /go-gin-example"]