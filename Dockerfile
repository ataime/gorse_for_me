# 使用官方的 Golang 镜像构建阶段
FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 复制 Gorse 源码
COPY . .

# 分阶段构建，减少每个阶段的内存消耗
RUN go mod tidy

RUN go build  --memory=4g  -o gorse cmd/gorse-in-one/main.go

# 使用一个更小的基础镜像
FROM debian:bullseye-slim

# 设置工作目录
WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/gorse /usr/local/bin/gorse

# 暴露服务端口
EXPOSE 8086
EXPOSE 8088

# 运行 Gorse
CMD ["gorse"]

