package postgres

import (
	"context"
	"search-keyword-service/internal/model"
	"search-keyword-service/internal/repository/db"
)

type KeywordRank struct {
	BaseULIDModel
	Keyword     string `gorm:"column:keyword;not null" json:"keyword"`
	Rank        int    `gorm:"column:rank" json:"rank"`
	Title       string `gorm:"column:title" json:"title"`
	Url         string `gorm:"column:url" json:"url"`
	Description string `gorm:"column:description" json:"description"`
}

type KeywordRankModelInterface interface {
	Create(ctx context.Context, in *KeywordRank) error
	Update(ctx context.Context, id string, value KeywordRank, field ...string) error
	FindOne(ctx context.Context, filter *model.KeywordRankFindQuery) (*KeywordRank, error)
	FindAll(ctx context.Context, filter *model.KeywordRankFindQuery) ([]KeywordRank, error)
}

var keywordRankRepo KeywordRankModelInterface = KeywordRankModel{}

type KeywordRankModel struct{}

func KeywordRankRepo() KeywordRankModelInterface {
	return keywordRankRepo
}

func (KeywordRankModel) Create(ctx context.Context, in *KeywordRank) error {
	if err := db.DBWithCtx(ctx).Create(in).Error; err != nil {
		return err
	}
	return nil
}

func (KeywordRankModel) Update(ctx context.Context, id string, value KeywordRank, field ...string) error {
	query := db.DBWithCtx(ctx).Model(&KeywordRank{BaseULIDModel: BaseULIDModel{ID: id}})
	if len(field) > 0 {
		query = query.Select(field)
	}

	if err := query.Updates(value).Error; err != nil {
		return err
	}
	return nil
}

func (KeywordRankModel) FindOne(ctx context.Context, filter *model.KeywordRankFindQuery) (*KeywordRank, error) {
	query := db.DBWithCtx(ctx).Model(&KeywordRank{}).Where(&KeywordRank{
		BaseULIDModel: BaseULIDModel{ID: filter.ID},
		Keyword:       filter.Keyword,
		Rank:          filter.Rank,
	})

	var res *KeywordRank
	if err := query.Take(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (KeywordRankModel) FindAll(ctx context.Context, filter *model.KeywordRankFindQuery) ([]KeywordRank, error) {
	var res []KeywordRank
	query := db.DBWithCtx(ctx).Model(&KeywordRank{}).Where(&KeywordRank{
		Keyword:     filter.Keyword,
		Rank:        filter.Rank,
		Description: filter.Description,
	})

	if err := query.Find(&res).Order("rank").Error; err != nil {
		return nil, err
	}
	return res, nil
}
