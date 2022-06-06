# build stage
FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN GOARCH=amd64 GOOS=linux go build -v -o cr-crboard .

# final stage
FROM alpine:latest
ENV GOPORT=8080
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/cr-crboard .
ENTRYPOINT ./cr-crboard