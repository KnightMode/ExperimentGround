# Build the Go API
FROM golang:alpine as builder
RUN apk add --no-cache ca-certificates
RUN apk update && apk add --no-cache git && apk add bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
ENV GO111MODULE=on
WORKDIR /src
COPY go.mod go.sum ./
ADD ./.profile.d ../app/.profile.d
RUN go mod download
COPY . .
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./src
CMD ["cd ../app/.profile.d/ && bash heroku-exec.sh", "./main"]