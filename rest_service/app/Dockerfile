FROM golang:latest
WORKDIR /build
COPY main.go .
RUN go build -o main main.go
CMD [". /main"]