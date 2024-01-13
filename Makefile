include .env
CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

createdb:
	docker exec -it ${DockerPostgresContainerName} createdb --username=${PostgresUser} --owner=${PostgresUser} ${PostgresDBName}

dropdb:
	docker exec -it ${DockerPostgresContainerName} dropdb --username=${PostgresUser} ${PostgresDBName}

psqlcontainer:
	docker run --name ${DockerPostgresContainerName} -d -p ${PostgresPort}:5432 -e POSTGRES_PASSWORD=${PostgresPassword} --env-file .env postgres:15-alpine3.16

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://${PostgresUser}:${PostgresPassword}@${PostgresHost}:${PostgresPort}/${PostgresDBName}?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://${PostgresUser}:${PostgresPassword}@${PostgresHost}:${PostgresPort}/${PostgresDBName}?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run ${APP_CMD_DIR}/main.go

.PHONE: start, createdb, dropdb, migration-up, migration-down, swag-init, psqlcontainer, rediscontainer
