FROM golang:1.22.2-alpine3.19
RUN apk add --no-cache gcc g++
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o bin main.go
CMD ["/app/bin"]