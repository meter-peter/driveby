.PHONY: build test vet clean up down validate integration-test docs-check docs-status thesis-status regen-openapi

# Top-level orchestration Makefile

build:
	cd driveby-cli && go build -ldflags "-X github.com/meter-peter/driveby/driveby-cli/internal/cli.version=$$(git describe --tags --always --dirty 2>/dev/null || echo dev) -X github.com/meter-peter/driveby/driveby-cli/internal/cli.commit=$$(git rev-parse --short HEAD 2>/dev/null || echo unknown) -X github.com/meter-peter/driveby/driveby-cli/internal/cli.date=$$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o driveby ./cmd/driveby

test:
	cd driveby-cli && go test -count=1 ./test/...

vet:
	cd driveby-cli && go vet ./...

clean:
	cd driveby-cli && rm -f driveby

# Docker targets
up:
	docker-compose up -d

down:
	docker-compose down

validate: build
	cd driveby-cli && ./driveby validate-only \
		--openapi ../apis/perfect-api/openapi.json \
		--host localhost \
		--port 8000

# Integration tests (requires docker-compose up)
integration-test: build up
	@echo "Waiting for services to be ready..."
	@sleep 3
	cd driveby-cli && go test -tags integration -count=1 -v ./test/...
	$(MAKE) down

# Regenerate OpenAPI spec from running perfect-api
regen-openapi: up
	@echo "Waiting for API to be ready..."
	@sleep 3
	curl -s http://localhost:8000/openapi.json | python3 -m json.tool > apis/perfect-api/openapi.json
	@echo "OpenAPI spec regenerated at apis/perfect-api/openapi.json"

# Helm targets
helm-template:
	helm template driveby kubernetes/helm/driveby/

helm-dry-run:
	helm template driveby kubernetes/helm/driveby/ | kubectl apply --dry-run=client -f -

# Documentation targets
docs-check:
	@bash tools/check-docs.sh

docs-status:
	@echo "=== Principle Implementation Status ==="
	@echo "ID    | Source File                | Tests             | Docs"
	@echo "------|----------------------------|-------------------|------------------"
	@for i in 1 2 3 4 5 6 7 8; do \
		id=$$(printf "P%03d" $$i); \
		id_lower=$$(printf "p%03d" $$i); \
		src="driveby-cli/internal/principles/$${id_lower}_*.go"; \
		tst="driveby-cli/test/$${id_lower}_test.go"; \
		doc="docs/principles/$${id}.md"; \
		src_status="MISSING"; tst_status="MISSING"; doc_status="MISSING"; \
		for f in $$src; do [ -f "$$f" ] && src_status="OK" && break; done; \
		[ -f "$$tst" ] && tst_status="OK"; \
		[ -f "$$doc" ] && doc_status="OK"; \
		printf "$$id  | %-26s | %-17s | %s\n" "$$src_status" "$$tst_status" "$$doc_status"; \
	done

thesis-status:
	@echo "=== Thesis Chapter Status ==="
	@for f in thesis/chapters/*.tex; do \
		lines=$$(wc -l < "$$f"); \
		basename_f=$$(basename "$$f"); \
		if [ "$$lines" -lt 20 ]; then \
			status="STUB"; \
		elif [ "$$lines" -lt 100 ]; then \
			status="DRAFT"; \
		else \
			status="WRITTEN"; \
		fi; \
		printf "%-30s %s (%d lines)\n" "$$basename_f" "$$status" "$$lines"; \
	done
