FROM golang:1.24.0 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/identity-server .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/identity-server /identity-server
EXPOSE 8080
ENTRYPOINT ["/identity-server"]