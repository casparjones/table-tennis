FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o tt-points

EXPOSE 8080

CMD ["/app/tt-points"]
