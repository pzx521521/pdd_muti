FROM golang
RUN mkdir /app
COPY . /app/
WORKDIR /app/
ENV CGO_ENABLED=0
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN go build /app/

FROM scratch
COPY --from=0 /app/pdd_muti .
CMD ["./pdd_muti"]
