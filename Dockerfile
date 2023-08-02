# build stage
FROM golang:1.20.7-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN GOARCH=amd64 GOOS=linux go build -v -o cr-portfolio-service .

# final stage
FROM alpine:latest
ENV GOPORT=8080
ENV MODE=production
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/cr-portfolio-service .
ENTRYPOINT ./cr-portfolio-service