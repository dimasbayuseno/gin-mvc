FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o gin-mvc

EXPOSE 8080

CMD ./gin-mvc