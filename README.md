# Search Keyword Contain In Url

## How to init, start service
1. Create a schema, table for database PostgreSQL
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

## Unit Test
1. Modify response handler keyword GetKeywordRank
```
	keyword := c.Param("word")
	// ctx := c.Request.Context()
	if len(keyword) == 0 {
		httputil.RespondWrapError(c, http.StatusBadRequest, "invalid_params", nil)
		return
	}

	// res := usecase.KeywordRankService().GetKeywordRank(ctx, keyword)
	res := []model.GetKeywordRankResponse{}
	res = append(res, model.GetKeywordRankResponse{
		Keyword: "qualgo",
		Rank:    1,
		Url:     "https://www.qualgo.io/",
		Title:   "Qualgo",
	},
		model.GetKeywordRankResponse{
			Keyword: "qualgo",
			Rank:    2,
			Url:     "/?FORM=Z9FD1",
			Title:   "",
		},
	)

	httputil.RespondWarpJSON(c, http.StatusOK, res)
```
2. Go to directory where you want to test
```
cd internal/handler/keyword
go test -v
```