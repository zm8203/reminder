# 基础镜像
FROM golang:1.17 AS builder

# 设置工作目录
WORKDIR /app

# 复制代码到容器中
COPY ./reminder .

# 编译应用程序
RUN cd /app && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o app .

# 第二阶段：构建最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 复制编译好的应用程序
COPY --from=builder /app/app .

# 暴露端口
EXPOSE 8080

# 启动应用程序
CMD ["./app"]
