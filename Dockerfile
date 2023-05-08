# Build
FROM golang:1.20-alpine3.17 as build
WORKDIR /build
COPY .  /build
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o govulnapi cmd/govulnapi/main.go

# Deploy
FROM alpine:3.17
WORKDIR /opt/govulnapi
COPY --from=build /build/govulnapi .
EXPOSE 8080 8081 8082

# Run
ENTRYPOINT ["/opt/govulnapi/govulnapi"]
