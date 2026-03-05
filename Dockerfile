# --- Build ---
FROM golang:1.25-alpine AS builder

# Added libc6-compat for glibc support
RUN apk add --no-cache gcc musl-dev curl libc6-compat

WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./

# Download dependency
RUN go mod download

# Copy the code
COPY . .

# Download the specific MUSL version for Alpine
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64-musl \
    && chmod +x tailwindcss-linux-x64-musl \
    && mv tailwindcss-linux-x64-musl /usr/local/bin/tailwindcss

# Compile Tailwind
RUN tailwindcss -i ./web/static/css/input.css -o ./web/static/css/style.css --minify

# Build Go
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