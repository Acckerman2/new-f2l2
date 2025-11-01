FROM golang:1.21-alpine3.18 AS builder
RUN apk update && apk upgrade --available && sync
WORKDIR /app

# 1. Copy ALL your code first (this is correct)
COPY . .

# 2. Run go mod tidy to clean up
RUN go mod tidy

# 3. THIS IS THE FIX:
# We will manually 'go get' the missing dependency
# that 'go mod tidy' is not finding.
RUN go get github.com/gotd/td/tg

# 4. Build the app (this will work now)
RUN CGO_ENABLED=0 go build -o /app/fsb -ldflags="-w -s" ./cmd/fsb

FROM scratch
COPY --from=builder /app/fsb /app/fsb
# Copy the SSL certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE ${PORT}
ENTRYPOINT ["/app/fsb", "run"]
