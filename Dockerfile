FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /webapp

FROM alpine:3.18
COPY --from=builder /webapp /webapp
EXPOSE 8080
ENTRYPOINT ["/webapp"]
