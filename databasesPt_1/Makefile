include .env
export

sql-main-run:
	@go run main.go

migrate-up:
	@migrate -path migrations -database $(CONN_STRING) up 1

migrate-down:
	@migrate -path migrations -database $(CONN_STRING) down 1

redis-run:
	@go run main_redis.go