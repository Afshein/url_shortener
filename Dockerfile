# Stage 1 — Build
FROM golang:1.25 AS builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server.exe ./cmd/web


# Stage 2 — Run (minimal)
FROM gcr.io/distroless/static-debian12 as runner
WORKDIR /app
COPY --from=builder /app/server.exe . 
CMD ["./server.exe -PORT=$PORT"]
