FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o api .

EXPOSE 8080

CMD ["/app/api"]