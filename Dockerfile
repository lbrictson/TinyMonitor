# Build in app in a Go container
FROM docker.io/golang:1.19.5-buster as builder
COPY ent /app/ent
COPY cmd /app/cmd
COPY pkg /app/pkg
COPY vendor /app/vendor
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
WORKDIR /app
RUN go test ./...
RUN go env -w GOPROXY=direct && go env -w GOSUMDB=off
RUN CGO_ENABLED=0 go build -o main cmd/tinymonitor/main.go
# Move artifact to smaller container with no Go tools installed
FROM mcr.microsoft.com/playwright:v1.30.0-focal
WORKDIR /app
COPY --from=builder /app/main tinymonitor
RUN SETUP_PLAYWRIGHT=true /app/tinymonitor server
ENTRYPOINT ["/app/tinymonitor", "server"]