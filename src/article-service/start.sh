#!/bin/bash
docker run --rm --name pg_article -d -p 5432:5432 -e POSTGRES_USER=article -e POSTGRES_PASSWORD=article -e POSTGRES_DB=article_db postgres:9.6
sleep 5
cd cmd
go run main.go