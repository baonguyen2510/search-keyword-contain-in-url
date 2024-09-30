package keyword

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type KeywordRank struct {
	Keyword string `json:"keyword"`
	Rank    int    `json:"rank"`
	URL     string `json:"url"`
	Title   string `json:"title"`
}

type KeywordRankResponse struct {
	Success bool          `json:"success"`
	Data    []KeywordRank `json:"data"`
}

// Test for GetKeywordHandler
func TestGetKeywordRank(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	keywordHandler := New()

	router.GET("/rank/:word", keywordHandler.GetKeywordRank)

	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, "/rank/qualgo", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Define the expected result
	expected := []KeywordRank{
		{
			Keyword: "qualgo",
			Rank:    1,
			URL:     "https://www.qualgo.io/",
			Title:   "Qualgo",
		},
		{
			Keyword: "qualgo",
			Rank:    2,
			URL:     "/?FORM=Z9FD1",
			Title:   "",
		},
	}

	var response KeywordRankResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expected, response.Data)
}
