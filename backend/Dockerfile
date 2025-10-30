FROM golang:1.25-alpine AS build
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/backend

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=0 /usr/local/bin/backend /usr/local/bin/backend
EXPOSE 8080
CMD ["/usr/local/bin/backend"]
