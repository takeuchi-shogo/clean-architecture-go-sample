up:
	docker-compose --env-file .env -f docker/docker-compose.yml up

rm:
	docker-compose --env-file .env -f docker/docker-compose.yml rm
