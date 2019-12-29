
FROM golang:1.10.1
WORKDIR /appdir
COPY . .
RUN go build -o app .
CMD ["./app"]
