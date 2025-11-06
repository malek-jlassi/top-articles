
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/top-articles ./cmd/app


FROM alpine:3.20

RUN adduser -D appuser
COPY --from=builder /bin/top-articles /usr/local/bin/top-articles
USER appuser

ENTRYPOINT ["top-articles"]
