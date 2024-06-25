# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd/
COPY docs ./docs/
COPY internal ./internal/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-gateway-svc ./cmd/auth-gateway-svc

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /auth-gateway-svc /auth-gateway-svc

EXPOSE 8080

ENTRYPOINT ["/auth-gateway-svc"]