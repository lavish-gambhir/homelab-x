FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o server .

FROM alpine:latest
COPY --from=builder /app/server /server

EXPOSE 8080
ENTRYPOINT [ "/server" ]
