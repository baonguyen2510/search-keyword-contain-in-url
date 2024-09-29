# Search Keyword Contain In Url

## How to init, start service
1. Create a schema, table for database
path: search-keyword-contain-in-url/deploy/migrate.sql
```
CREATE SCHEMA IF NOT exists "search";

CREATE TABLE IF NOT EXISTS search.keyword_rank ( 
	id TEXT NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	keyword TEXT,
	rank BIGINT DEFAULT 0,
	title text,
	url text,
	description text,
	PRIMARY KEY(id)
);
```
2. cd search-keyword-contain-in-url/ 
3. go mod tidy

## How to run service
1. Fill environment variables
/search-keyword-contain-in-url/configs/config.yaml
```
DB_CONNECTION:
DB_HOST: localhost
DB_PORT: 5432
DB_USERNAME: postgres
DB_PASSWORD: "123456789"
DB_NAME: postgres
DB_SCHEMA: search
DB_LOGSQL: true
DB_AUTOMIGRATE: false

REDIS_SINGLE: true
REDIS_ADDR: localhost:6379
REDIS_PASSWORD: ""

BASIC_AUTH_USER: admin
BASIC_AUTH_PASSWORD: password

CONFIG_TIME_SCHEDULE: 3600

SEARCH_ENGINE_ADDR: "google"
```


## Curl cmd for call API
1. API for health check service
```
curl --location 'localhost:7003/v1/search/health'
```

2. API for get list rank of keyword
```
curl --location 'localhost:7003/v1/search/keyword/rank/:word' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='
```

3. API for update rank of keyword
```
curl --location --request POST 'localhost:7003/v1/search/keyword/sync/qualgo' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='
```

## How to generate, access swagger
1. Install swagger latest version
```
go get -u github.com/swaggo/swag/cmd/swag
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
```
2. Update description in handler of API
3. Generate swagger
```
swag init
```
4. Access swagger
```
http://localhost:7003/v1/search/swagger/index.html
```