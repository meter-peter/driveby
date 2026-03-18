#!/usr/bin/env bash
set -euo pipefail

echo "=== DriveBy Local Demo ==="

# Start services
echo "Starting perfect-api..."
docker-compose up -d
sleep 5

# Wait for healthy
echo "Waiting for API to be ready..."
for i in $(seq 1 30); do
  if curl -sf http://localhost:8000/test/health > /dev/null 2>&1; then
    echo "API is ready!"
    break
  fi
  sleep 1
done

# Run all modes
echo ""
echo "=== Running validate-only ==="
cd driveby-cli && ./driveby validate-only \
  --openapi ../apis/perfect-api/openapi.json \
  --host localhost --port 8000 \
  --validation-mode minimal \
  --report-dir /tmp/driveby-demo

echo ""
echo "=== Running function-only ==="
./driveby function-only \
  --openapi ../apis/perfect-api/openapi.json \
  --host localhost --port 8000 \
  --report-dir /tmp/driveby-demo

echo ""
echo "=== Running load-only ==="
./driveby load-only \
  --openapi ../apis/perfect-api/openapi.json \
  --host localhost --port 8000 \
  --concurrent-users 5 --test-duration 10s \
  --report-dir /tmp/driveby-demo

echo ""
echo "=== Demo complete! Reports in /tmp/driveby-demo ==="
ls -la /tmp/driveby-demo/
