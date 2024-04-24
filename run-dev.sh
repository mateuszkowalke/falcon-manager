#!/usr/bin/env bash

docker compose -f docker-compose.db.yml up -d
while ! timeout 1 sh -c 'nc -zv localhost 5432' &>/dev/null; do : ; done
export JWT_SECRET=somesecret
export DATABASE_URL=postgresql://testuser:testpass@localhost:5432/falcon-manager?sslmode=disable
air go run
docker compose down
