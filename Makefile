docker-rebuild:
	docker-compose -f deployments/local/docker-compose.yaml build --no-cache
docker-up:
	docker-compose -f deployments/local/docker-compose.yaml up -d
docker-down:
	docker-compose -f deployments/local/docker-compose.yaml down
docker-status:
	docker-compose -f deployments/local/docker-compose.yaml ps -a
migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:local@localhost:55432/?sslmode=disable up
migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:local@localhost:55432/?sslmode=disable down --all
down: docker-down
up: docker-up