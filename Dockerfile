FROM golang:1.21 AS builder
WORKDIR /module
COPY .  .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /module/app .
EXPOSE 8000
CMD ["./app"]