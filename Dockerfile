FROM golang
ENV GO111MODULE=on

WORKDIR /app

COPY . /app
RUN go build

EXPOSE 9090

ENTRYPOINT ["/app/hello-world"]
