FROM golang:1.13.4

COPY . /app
WORKDIR /app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app

CMD ["./app"]
