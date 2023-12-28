FROM golang

WORKDIR /app
COPY . .

RUN go mod download
EXPOSE 8081

CMD ["go", "run", "cmd/main.go"]

# FROM golang AS builder
# WORKDIR /build
# COPY . .
# RUN go clean --modcache
# RUN go build -mod=readonly -o store cmd/main.go

# FROM alpine

# WORKDIR /build
# COPY --from=builder /build/store /
# EXPOSE 8081
# CMD ["/store"]