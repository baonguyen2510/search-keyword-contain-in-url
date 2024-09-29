package keyword

import (
	"context"
	"net/http"
	"search-keyword-service/internal/usecase"
	"search-keyword-service/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

// GetKeywordRank godoc
// @Summary Retrieve keyword rank
// @Description Get the current rank of a specified keyword.
// @Tags GetKeywordRank
// @Produce json
// @Param word query string true "Keyword to search" true
// @Success 200 {object} model.GetKeywordRankResponse "Success"
// @Failure 400 {object} model.KeywordResponse "Failed"
// @Router /keyword/rank/:word [get]
func (Handler) GetKeywordRank(c *gin.Context) {

	keyword := c.Param("keyword")
	ctx := c.Request.Context()
	if len(keyword) == 0 {
		httputil.RespondWrapError(c, http.StatusBadRequest, "invalid_params", nil)
		return
	}

	res := usecase.KeywordRankService().GetKeywordRank(ctx, keyword)

	httputil.RespondWarpJSON(c, http.StatusOK, res)
}

// SyncKeywordRank godoc
// @Summary Update keyword rank
// @Description Update rank of a specified keyword.
// @Tags SyncKeywordRank
// @Produce json
// @Param word query string true "Keyword to update" true
// @Success 200 {object} model.KeywordResponse "Success"
// @Failure 400 {object} model.KeywordResponse "Failed"
// @Router /keyword/rank/:word [post]
func (Handler) SyncKeywordRank(c *gin.Context) {

	keyword := c.Param("word")
	if len(keyword) == 0 {
		httputil.RespondWrapError(c, http.StatusBadRequest, "invalid_params", nil)
		return
	}

	ctx := context.WithoutCancel(c.Request.Context())
	// Start background task to update keyword ranks
	go func(ctx context.Context, keyword string) {
		err := usecase.KeywordRankService().SyncKeywordRank(ctx, keyword)
		if err != nil {
			return
		}
	}(ctx, keyword)

	httputil.RespondWarpJSON(c, http.StatusOK, nil)
}

func (Handler) SyncAllKeywordsRank() {

	ctx := context.Background()
	usecase.KeywordRankService().SyncAllKeywordsRank(ctx)
}
