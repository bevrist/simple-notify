FROM golang:1

ENV CGO_ENABLED=1

WORKDIR /app

COPY . .

RUN go get ./...
RUN go build -o core .

ENTRYPOINT ./core
