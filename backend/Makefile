# Run from this directory (`make -C backend` from the repo root works).
# Paths stay relative. Make's dir/patsubst helpers split on spaces, and this
# checkout lives in a directory whose name contains one.
SHELL := /bin/sh

SERVICES := api-gateway identity billing membership passport metering payments documents calendar history media notification content
DB_SERVICES := identity billing membership passport metering payments documents calendar history media notification content
MIGRATE_IMAGE := migrate/migrate:v4.18.2
NETWORK := kvartirnik

PING_TARGETS := \
	api-gateway:8080 \
	identity:8081 \
	billing:8082 \
	membership:8083 \
	passport:8084 \
	metering:8085 \
	payments:8086 \
	documents:8087 \
	calendar:8088 \
	history:8089 \
	media:8090 \
	notification:8091 \
	content:8092

.PHONY: build up down logs migrate ping

build:
	@set -eu; \
	test -f go.mod; \
	mkdir -p bin; \
	for s in $(SERVICES); do \
		echo "==> $$s"; \
		go build -o "bin/$$s" "./cmd/$$s"; \
	done

up:
	docker compose -p kvartirnik up --build -d

down:
	docker compose -p kvartirnik down

logs:
	docker compose -p kvartirnik logs -f --tail=200

migrate:
	@set -eu; \
	test -f go.mod; \
	for s in $(DB_SERVICES); do \
		echo "==> migrate $$s"; \
		docker run --rm --network $(NETWORK) \
			-v "$(CURDIR)/migrations/$$s:/migrations:ro" \
			$(MIGRATE_IMAGE) \
			-path=/migrations \
			-database "postgres://$$s:$$s-draft@$$s-db:5432/$$s?sslmode=disable" \
			up; \
	done

ping:
	@set -eu; \
	fail=0; \
	for item in $(PING_TARGETS); do \
		name=$${item%%:*}; \
		port=$${item##*:}; \
		if ! body=$$(curl -fsS "http://127.0.0.1:$$port/ping"); then \
			echo "FAIL $$name"; \
			fail=1; \
			continue; \
		fi; \
		printf '%s' "$$body" | grep -q "\"service\":\"$$name\"" || { echo "FAIL $$name unexpected $$body"; fail=1; continue; }; \
		printf '%s' "$$body" | grep -q "\"status\":\"ok\"" || { echo "FAIL $$name unexpected $$body"; fail=1; continue; }; \
		echo "ok $$name $$body"; \
	done; \
	exit $$fail
