FROM golang as builder
WORKDIR /app/
COPY . .
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV CGO_ENABLED=0
RUN go build .

FROM busybox as runner
WORKDIR /app/
COPY --from=builder /app/ /app/
CMD ["./pdd_muti"]
