# api-gateway draft image. Build context is the backend module root:
#   docker build -f cmd/api-gateway/Dockerfile .
FROM golang:1.27.1-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal/ ./internal/
COPY cmd/api-gateway/ ./cmd/api-gateway/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/service ./cmd/api-gateway

FROM scratch
COPY --from=build /out/service /service
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/service"]
