FROM golang AS builder

WORKDIR /usr/local/src

# Dependencies
COPY go.mod ./ 
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/store cmd/main.go

FROM alpine as runner

COPY --from=builder usr/local/src/bin/store /
COPY --from=builder usr/local/src/configs /configs


CMD ["/store"]