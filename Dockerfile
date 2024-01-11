
FROM golang:1.18.1 AS build
WORKDIR /app

COPY main.go /

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY * ./

RUN CGO_ENABLED=1 GOOS=linux 
RUN go build -o /docker-gs-ping


EXPOSE 8080


ENTRYPOINT ["/docker-gs-ping"]