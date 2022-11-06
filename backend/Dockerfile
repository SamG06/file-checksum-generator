FROM golang:1.19.3-alpine

WORKDIR /app

COPY *.go ./

RUN go build server.go

EXPOSE 3000

CMD ["/app/server"]