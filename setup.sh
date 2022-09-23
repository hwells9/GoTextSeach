#!/bin/bash

docker-compose up -d

go run main.go -executeBuildmodels true

go run main.go -executeDataMigration true
