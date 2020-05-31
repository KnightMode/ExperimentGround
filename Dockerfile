# Build the Go API
FROM golang:alpine as builder
RUN apk add --no-cache ca-certificates
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
WORKDIR /src
COPY go.mod go.sum ./

RUN go mod download

COPY . .
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./src
CMD ["./main"]