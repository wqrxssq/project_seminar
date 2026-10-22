# Backend

Draft Go services for Квартирник. Domain boundaries are in [`../docs/c4.md`](../docs/c4.md).

Each process is its own binary and, except `api-gateway`, its own Postgres. The phone talks only to `api-gateway`. This tree does not implement that routing yet.

Module path: `github.com/wqrxssq/project_seminar/backend`. The module asks for Go 1.22 or newer. Images build with Go 1.27.1.

## Layout

```
backend/
  api/openapi.yaml          draft external HTTP contract
  cmd/<service>/            one main and one Dockerfile per process
  deploy/nats.conf          JetStream cluster config shared by the three nodes
  internal/httpserver/      GET /ping and graceful shutdown
  migrations/<service>/     first golang-migrate revision, app_meta only
  docker-compose.yml        project name kvartirnik
  Makefile
```

Processes: `api-gateway`, `identity`, `billing`, `membership`, `passport`, `metering`, `payments`, `documents`, `calendar`, `history`, `media`, `notification`, `content`.

## Run

From this directory:

```bash
make build    # compile every binary into bin/
make up       # build images and start the stack in the background
make migrate  # apply SQL to each domain database (stack must be up)
make ping     # curl every GET /ping
make logs     # follow compose logs
make down     # stop and remove containers; volumes stay
```

`make build` needs a local Go toolchain and does not need Docker. `make up`, `make migrate`, and `make down` need Docker.

`make migrate` runs `migrate/migrate:v4.18.2` on the compose network `kvartirnik`. The host does not need the migrate CLI. Postgres ports are not published.

Wipe draft data with:

```bash
docker compose -p kvartirnik down -v
```

### What answers

`GET /ping` returns `200` and `{"service":"<name>","status":"ok"}`. The listen address is `LISTEN_ADDR`, default `:8080`. SIGINT and SIGTERM shut the process down.

The binaries record whether `DATABASE_URL`, `NATS_URL`, and `S3_ENDPOINT` are present. They do not open those connections.

Host ports below are for this draft so `make ping` can reach every process. The product ingress is `api-gateway`.

| Process | Host port |
|---|---|
| api-gateway | 8080 |
| identity | 8081 |
| billing | 8082 |
| membership | 8083 |
| passport | 8084 |
| metering | 8085 |
| payments | 8086 |
| documents | 8087 |
| calendar | 8088 |
| history | 8089 |
| media | 8090 |
| notification | 8091 |
| content | 8092 |

MinIO, the draft stand-in for private S3, is on `localhost:9000` (API) and `localhost:9001` (console). Root user `kvartirnik`, password `kvartirnik-draft`. The bucket name the media service expects is `media`. Compose does not create that bucket.

NATS JetStream is three nodes (`nats-1`, `nats-2`, `nats-3`) on the compose network. Client and route ports are not published on the host. The draft cluster has no auth.

Each domain database is its own Postgres 16 container. User, password, and database name are the service name, and the password is `<service>-draft` (for example `identity` / `identity-draft`). Those values are local draft credentials.

`api-gateway` has no database and no `DATABASE_URL`. `content` has Postgres. NATS is wired for the services that publish or consume the facts in the architecture note: `identity`, `passport`, `metering`, `payments`, `documents`, `calendar`, `history`, `media`, `notification`.

## OpenAPI

[`api/openapi.yaml`](api/openapi.yaml) is a draft of the external HTTP API: passport, readings, payments, documents, invites, premium, pushes, and stories. Handlers for those paths are not registered. The file points back at `docs/c4.md`, `docs/PRD.md`, and `docs/SRS.md`.

## Left for a later pass

- Domain HTTP handlers and the gRPC calls between gateway and domains
- JWT checks and `Idempotency-Key` handling
- River workers, outbox, and inbox (the queue stays in each domain Postgres when it exists)
- Domain tables beyond the `app_meta` marker
- PgBouncer, a Postgres standby, and Kubernetes
- SMS, APNs, FCM, App Store, and Google Play clients
- Presigned upload and creating the `media` bucket
- NATS authentication and TLS
- RabbitMQ and Redis
