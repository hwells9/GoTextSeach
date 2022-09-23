GoTextSeach
===========
**A simple text search using postgresql text search feature**

Setup
-----
- Run *bash setup.sh* in root directory to create and start db container,
    created tables, and migrate data from marvel api to the database
- Run *base remove.sh* to stop and remove the container **Will delete all database data**

Docker
------
- run command *docker compose up* inside of the root directory to create database

Data Migration
--------------
- The data migration is built into the main.go file.
- Run *go run main.go -executeDataMigration True* in order to migrate data.

Database
----------
- PostgreSQL 14.5 
- ERD created on: draw.io
- To edit or update erd, open existing diagram *mc_text_search_erd.drawio* on draw.io.
- Screen shot of erd also located in docs

Envrionment Variables
---------------------
- Create *env-vars.env* file in root directory
- Add the following variables and their values if not present:
    - PRIVATE_API_KEY=
    - PUBLIC_API_KEY=
    - DATABASE_NAME=
    - DATABASE_ADDRESS=
    - DATABASE_USERNAME=
    - DATABASE_PASSWORD=
- They will automatically be pulled in where needed
- Import "github.com/spf13/viper" to access env vars