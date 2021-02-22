FROM golang:latest

WORKDIR /app

COPY ./Post-Service .

RUN go mod download

RUN go build .

CMD ["./postservice"]