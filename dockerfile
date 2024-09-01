FROM golang:1.23.0

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o bucket-app ./main.go
