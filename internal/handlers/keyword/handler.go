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

	httputil.RespondWarpJSON(c, http.StatusOK, res)
}

func (Handler) SyncKeywordRank(c *gin.Context) {

	keyword := c.Param("word")
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
