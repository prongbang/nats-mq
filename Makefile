start:
	docker-compose up --build -d
clear:
	docker system prune -a
run:
	go run main.go