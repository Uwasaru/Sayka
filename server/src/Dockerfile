## Build
FROM golang:latest AS build

WORKDIR /go/src/github.com/Uwasaru/Sayka

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

## Deploy
# FROM gcr.io/distroless/base-debian10

