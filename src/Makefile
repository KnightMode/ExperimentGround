start:
	@echo "Starting database...."
	docker-compose -f ../docker-compose.yml up -d

stop:
	@echo "Stopping database, removing networks........."
	docker-compose down
logs:
	docker-compose -f ../docker-compose.yml logs -f -t

build:
	go build -o ./web && ./web