# Build stage
FROM golang:1.21-alpine3.18 AS builder
RUN apk update && apk upgrade --available && sync
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /app/fsb -ldflags="-w -s" ./cmd/fsb

# Run stage
FROM scratch
COPY --from=builder /app/fsb /app/fsb
EXPOSE 8080  # ✅ Render expects your app to listen on port 8080
ENTRYPOINT ["/app/fsb", "run"]
