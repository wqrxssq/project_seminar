# media draft image. Build context is the backend module root:
#   docker build -f cmd/media/Dockerfile .
FROM golang:1.27.1-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal/ ./internal/
COPY cmd/media/ ./cmd/media/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/service ./cmd/media

FROM scratch
COPY --from=build /out/service /service
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/service"]
