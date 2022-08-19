# syntax=docker/dockerfile:1

FROM golang:1.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download
RUN go mod verify
RUN go build -o /docker-gorest

EXPOSE 8080

CMD [ "/docker-gorest" ]