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

func (Handler) GetKeywordRank(c *gin.Context) {

	keyword := c.Param("keyword")
	ctx := c.Request.Context()
	res := usecase.KeywordRankService().GetKeywordRank(ctx, keyword)
	if len(res) == 0 {
		httputil.RespondWarpJSON(c, http.StatusOK, "Keyword not found")
		return
	}

	httputil.RespondWarpJSON(c, http.StatusOK, res)
}

func (Handler) SyncKeywordRank(c *gin.Context) {

	keyword := c.Param("word")
	ctx := c.Request.Context()
	err := usecase.KeywordRankService().SyncKeywordRank(ctx, keyword)
	if err != nil {
		httputil.RespondWrapError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httputil.RespondWarpJSON(c, http.StatusOK, nil)
}

func (Handler) SyncAllKeywordsRank() {

	ctx := context.Background()
	usecase.KeywordRankService().SyncAllKeywordsRank(ctx)
}
