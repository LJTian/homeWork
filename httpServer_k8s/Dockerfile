FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 移动到工作目录：/build
WORKDIR /build

# 将我们的代码编译成二进制可执行文件 app
COPY . .
RUN go mod tidy
RUN go mod download
# 将代码复制到容器中
RUN go build -o bin/httpserver src/main.go

###################
# 接下来创建一个小镜像
###################
FROM alpine:3.13

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY ./etc ./httpServer/etc
COPY ./shell ./httpServer/shell
COPY ./static ./httpServer/static
COPY ./log  ./httpServer/log
COPY --from=builder /build/bin/ ./httpServer/bin/

#外放端口
EXPOSE 8009

# 需要运行的命令
ENTRYPOINT ["./httpServer/bin/httpserver"]

