#源镜像
#FROM golang:latest
FROM alpine:latest
#作者
MAINTAINER "huyang@udglin.com"
#设置工作目录
WORKDIR $GOPATH/src/Hoo/DevOps
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/Hoo/DevOps
#go构建可执行文件
#RUN go build .
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["./main"]
