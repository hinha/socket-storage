FROM golang:1.14.10-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc g++ libc-dev

RUN apk add build-base


# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

LABEL maintainer="Martinus <martinuz.dawan9@gmail.com>"

WORKDIR /app

ADD . /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Build the Go app
RUN go build -o main ./cmd/s3/main.go

EXPOSE 8082
