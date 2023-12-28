FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod download
EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]

# FROM golang AS builder
# WORKDIR /build

# COPY . .
# RUN go clean --modcache
# RUN go build -mod=readonly -o auth ./cmd/main.go

# FROM alpine

# COPY --from=builder /build/auth /build

# CMD ["/build"]

