# Simple multi-stage Dockerfile for the Top Articles Go app
# Note: Ensure your go.mod has a valid Go version like: `go 1.22`
# (Patch versions like 1.25.4 are invalid in the `go` directive.)

# ---- Build stage ----
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Pre-cache dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a static binary for Linux (no CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/top-articles ./cmd/app

# ---- Runtime stage ----
FROM alpine:3.20
# Create non-root user
RUN adduser -D appuser
COPY --from=builder /bin/top-articles /usr/local/bin/top-articles
USER appuser

# Default command
ENTRYPOINT ["top-articles"]
