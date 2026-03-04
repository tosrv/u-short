# --- Build ---
FROM golang:1.21-alpine AS builder

# Install requirement tools SQLite & Tailwind Standalone
RUN apk add --no-cache gcc musl-dev curl

WORKDIR /app
COPY . .

# 1. Download & compile Tailwind CLI
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x tailwindcss-linux-x64 \
    && ./tailwindcss-linux-x64 -i ./web/static/css/input.css -o ./web/static/css/style.css --minify

# 2. Build Go (CGO_ENABLED=1 for SQLite)
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/server/main.go

# --- Runtime ---
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary & folder web (templates & static)
COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

# Create folder data for Database SQLite (Volume)
RUN mkdir -p /data

# Set ENV Database URL & Port
ENV DATABASE_URL=/data/u_short.db
ENV PORT=8080

EXPOSE 8080
CMD ["./main"]