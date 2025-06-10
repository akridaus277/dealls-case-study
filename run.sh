#!/bin/bash

set -e

# Daftar nama folder service
SERVICES=(
  "api-gateway-service"
  "employee-service"
  "authentication-service"
  "attendance-service"
  "overtime-service"
  "reimbursement-service"
  "payroll-service"
  "audit-service"
)

echo ""
echo "🚀 Menjalankan semua test..."

for service in "${SERVICES[@]}"; do
  echo "🔍 Testing di $service..."
  (cd "$service" && go test ./...)
done

echo ""
echo "✅ Semua test berhasil!"
echo "🚀 Menjalankan semua service..."

for service in "${SERVICES[@]}"; do
  echo "▶️ Menjalankan $service..."
  (cd "$service" && go run ./cmd/main.go) &
done

echo ""
echo "🌐 Semua service sedang berjalan..."

wait
