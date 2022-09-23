#!/bin/bash

# Create container and the database
docker compose up -d

# let database intialize
sleep 5

# Builds all the tables in the database
go run main.go -executeBuildTables true

# Migrates data from the marvel api into our database
go run main.go -executeDataMigration true
