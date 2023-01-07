FROM golang:latest

COPY ./ ./
RUN go build -o ad-service ./cmd/ad-service/main.go
CMD ["./ad-service"]