.PHONY: consumer

test:
	docker-compose up --build

consumer:
	docker build --target builder --tag go-kafka-consumer -f Dockerfile ./consumer
	docker run go-kafka-consumer

producer:
	docker build --target builder --tag go-kafka-producer -f Dockerfile ./producer
	docker run go-kafka-producer
