FROM golang:1.23

WORKDIR /src

COPY . .

RUN go mod download
RUN go build ./cmd/nsc
