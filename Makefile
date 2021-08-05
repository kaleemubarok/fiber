#!make
include .env

#Docker
docker.network:
	docker network create -d bridge dev-network



#Database
docker.db.run:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgress \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgress \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres



#Migrate
migrate.up:
	migrate \
		-path $(PWD)/platform/migrations \
		-database "${MIGRATE_DB}" \
		up

migrate.down:
	migrate \
		-path $(PWD)/platform/migrations \
		-database "${MIGRATE_DB}" \
		down

