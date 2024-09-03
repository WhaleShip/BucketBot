FROM golang:1.23.0

RUN go version
ENV GOPATH=/

COPY ./ ./

EXPOSE 8080

RUN go mod download
RUN go build -o bucket-app ./cmd/app/main.go
