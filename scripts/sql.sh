#!/bin/bash

docker compose -f ./docker/docker-compose.db.yaml -f ./docker/docker-compose.local.yaml exec db bash -c "mysql -u docker -pdocker cateiru"
