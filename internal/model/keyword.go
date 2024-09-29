package model

type GetKeywordRankInfoQuery struct {
	Keyword string `form:"keyword" json:"keyword"`
}

type KeywordRankFindQuery struct {
	ID          string
	Keyword     string
	Rank        int
	Description string
}

type SyncKeywordRankResponse struct {
	Keyword string                   `json:"keyword"`
	Ranks   []map[string]interface{} `json:"ranks"`
}

type GetKeywordRankResponse struct {
	Keyword string `json:"keyword"`
	Rank    int    `json:"rank"`
	Url     string `json:"url"`
	Title   string `json:"title"`
}
