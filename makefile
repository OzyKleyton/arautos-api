start:
	docker-compose exec api go run cmd/server/main.go

up:
	docker-compose up -d

down:
	docker-compose down --timeout 0