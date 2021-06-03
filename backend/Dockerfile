FROM golang:1.16.0-stretch as builder
COPY . /server
WORKDIR /server
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go 

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /server .
EXPOSE 80
CMD ["./server"]
