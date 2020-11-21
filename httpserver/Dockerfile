FROM golang:1.14.3-alpine
LABEL maintainer="tzz1002@gmail.com"
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 8080
ENTRYPOINT ["/app/httpserver"]