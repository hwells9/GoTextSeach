GoTextSeach
===========
**A simple text search using postgresql text search feature**

Setup
-----
- Run `*bash setup.sh*` in root directory to create and start db container,
    created tables, and migrate data from marvel api to the database
- Run `*bash remove.sh*` to stop and remove the container **Warning: Will delete all database data**
- Run `*go run .*` from the root directory to start api

Docker
------
** If you do not run setup first and want to create container **
- run command `*docker compose up -d*` inside of the root directory to create database

Postman
-------
- In the `*docs/postman*` directory there is the environment and collection
  json files
- Import them into postman
- **Will Have to create a token and populate variable in environment first**

Authentication
--------------
- Once postman is setup, make the `*Register User*` call in the User folder
- Then use the `*Request Token*` call in the Authentication folder
- Copy the token from the response, and paste into the environment `*token*` variable
- **If the body of the `*Register User*` call is modified, update the body in the `*Request Token*` accordingly**

Data Migration
--------------
- The data migration is built into the main.go file.
- Run `*go run main.go -executeDataMigration True*` in order to migrate data.

Database
----------
- PostgreSQL 14.5 
- ERD created on: draw.io
- To edit or update erd, open existing diagram `*mc_text_search_erd.drawio*` on draw.io.
- Screen shot of erd also located in docs

Envrionment Variables
---------------------
- Create `*env-vars.env*` file in root directory
- Add the following variables and their values if not present:
    - PRIVATE_API_KEY=
    - PUBLIC_API_KEY=
    - DATABASE_NAME=
    - DATABASE_ADDRESS=
    - DATABASE_USERNAME=
    - DATABASE_PASSWORD=
- They will automatically be pulled in where needed
- Import "github.com/spf13/viper" to access env vars

