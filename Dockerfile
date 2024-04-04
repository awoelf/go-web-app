FROM golang:1.22.2-alpine3.19
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ls
RUN go mod tidy
RUN go build -o bin ./cmd/server/main.go
CMD ["/app/main"]