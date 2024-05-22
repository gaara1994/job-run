# 使用golang镜像作为基础镜像
FROM golang
# 将这段代码复制到镜像中的/app目录
ADD main /app/main
# 设置工作目录为/app
WORKDIR /app
# 设置ENTRYPOINT和CMD，允许在运行时传递命令行参数
ENTRYPOINT ["/app/main"]
CMD ["--result=succeed", "--second=10"]
