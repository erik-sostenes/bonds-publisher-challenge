#!/bin/bash

CONTAINER_NAME=$(docker-compose ps -q postgres)

SQL_FILE="../sql/inserts.sql"

docker exec -i $CONTAINER_NAME psql -U postgres -d bonds_db -f $SQL_FILE
