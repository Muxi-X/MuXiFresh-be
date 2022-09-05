FROM golang:1.16
RUN mkdir /app 
ADD . /app
RUN go env -w GOPROXY="https://goproxy.cn,direct"
WORKDIR /app
RUN make 
CMD ["./main"]

