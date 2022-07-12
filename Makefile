migrateup:
	migrate -path migrations -database "postgresql://testuser:testpass@localhost:5432/falcon_manager?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://testuser:testpass@localhost:5432/falcon_manager?sslmode=disable" -verbose down

seed:
	docker exec falcon-manager-pgdb-1 psql "postgresql://testuser:testpass@localhost:5432/falcon_manager" -c "$(cat seed.sql)"

.PHONY: migrateup migratedown seed
