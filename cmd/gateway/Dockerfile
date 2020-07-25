FROM golang:latest
COPY . /go/src/go_test
WORKDIR /go/src/go_test
RUN ls -la
RUN go mod download
RUN go build -o main cmd/gateway/main.go
EXPOSE 8081
CMD ["./main"]

