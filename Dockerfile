FROM golang:1.21-alpine3.18 AS builder
RUN apk update && apk upgrade --available && sync
WORKDIR /app

# 1. Copy *only* the module files first
COPY go.mod go.sum ./

# 2. Download dependencies and run go mod tidy to fix inconsistencies
# This is the new, important step
RUN go mod download
RUN go mod tidy

# 3. Now copy the rest of your code
COPY . .

# 4. Build the app (this should work now)
RUN CGO_ENABLED=0 go build -o /app/fsb -ldflags="-w -s" ./cmd/fsb

FROM scratch
COPY --from=builder /app/fsb /app/fsb
# Copy the SSL certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE ${PORT}
ENTRYPOINT ["/app/fsb", "run"]
