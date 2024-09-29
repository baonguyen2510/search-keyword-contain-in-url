package usecase

import (
	"context"
	"search-keyword-service/common"
	"search-keyword-service/internal/model"
	"search-keyword-service/internal/repository/postgres"
	"search-keyword-service/pkg/log"
)

type KeywordRankServiceInterface interface {
	SyncKeywordRank(ctx context.Context, keyword string) error
	SyncAllKeywordsRank(ctx context.Context)
	GetKeywordRank(ctx context.Context, keyword string) []model.GetKeywordRankResponse
}

var keywordRankService KeywordRankServiceInterface = &KeywordRankServiceImpl{}

type KeywordRankServiceImpl struct {
}

func KeywordRankService() KeywordRankServiceInterface {
	return keywordRankService
}

func (kr KeywordRankServiceImpl) SyncKeywordRank(ctx context.Context, keyword string) error {

	res, err := postgres.KeywordRankRepo().FindAll(ctx, &model.KeywordRankFindQuery{
		Keyword:     keyword,
		Description: common.Qualify,
	})
	if err != nil {
		return err
	}

	ranks, err := SearchEngineService().SearchKeyword(ctx, keyword)
	if err != nil {
		return err
	}

	if len(res) > 0 {
		// Update ranking when sync keyword
		m := make(map[string]postgres.KeywordRank)
		for _, v := range res {
			m[v.Url] = v
		}

		for _, v := range ranks {
			rank := v["rank"]
			rankInt, ok := rank.(int)
			if !ok {
				log.Errorf("Error type assertion failed for rank: %v", rank)
				continue
			}

			linkUrl := v["url"].(string)
			if m[linkUrl].Url == v["url"].(string) && m[linkUrl].Rank != rankInt && m[linkUrl].Rank > rankInt {
				err = postgres.KeywordRankRepo().Update(ctx, m[linkUrl].ID, postgres.KeywordRank{
					Rank: rankInt,
				}, "rank")
				if err != nil {
					log.Errorf("postgres.KeywordRankRepo().Update: err %v", err)
					continue
				}
			}
		}
	} else {
		for _, v := range ranks {
			rank := v["rank"]
			rankInt, ok := rank.(int)
			if !ok {
				log.Errorf("Error type assertion failed for rank: %v", rank)
				continue
			}

			err = postgres.KeywordRankRepo().Create(ctx, &postgres.KeywordRank{
				Keyword:     keyword,
				Rank:        rankInt,
				Title:       v["title"].(string),
				Url:         v["url"].(string),
				Description: v["qualified"].(string),
			})
			if err != nil {
				log.Errorf("Error creating keyword rank: %v", err)
				continue
			}
		}
	}

	return nil
}

func (kr KeywordRankServiceImpl) SyncAllKeywordsRank(ctx context.Context) {
	log.Infof("SyncAllKeywordsRank starting...")
	res, err := postgres.KeywordRankRepo().FindAll(ctx, &model.KeywordRankFindQuery{
		Description: common.Qualify,
	})
	if err != nil {
		log.Errorf("postgres.KeywordRankRepo().FindAll: err %v", err)
		return
	}

	m := make(map[string]postgres.KeywordRank)
	for _, v := range res {
		m[v.Url] = v
	}

	for _, v := range res {
		ranks, err := SearchEngineService().SearchKeyword(ctx, v.Keyword)
		if err != nil {
			log.Warnf("SearchEngineService().SearchKeyword: %v", err)
			continue
		}

		for _, i := range ranks {
			rank := i["rank"]
			rankInt, ok := rank.(int)
			if !ok {
				log.Errorf("Error type assertion failed for rank: %v", rank)
				continue
			}

			linkUrl := i["url"].(string)
			if m[linkUrl].Url == i["url"].(string) && m[linkUrl].Rank != rankInt && m[linkUrl].Rank > rankInt {
				err = postgres.KeywordRankRepo().Update(ctx, m[linkUrl].ID, postgres.KeywordRank{
					Rank: rankInt,
				})
				if err != nil {
					log.Warnf("postgres.KeywordRankRepo().Update: err %v", err)
					continue
				}
			}
		}
		log.Infof("SyncAllKeywordsRank: %+v success", v)
	}
	log.Infof("SyncAllKeywordsRank end...")
}

func (kr KeywordRankServiceImpl) GetKeywordRank(ctx context.Context, keyword string) []model.GetKeywordRankResponse {
	resQualify, err := postgres.KeywordRankRepo().FindAll(ctx, &model.KeywordRankFindQuery{
		Keyword:     keyword,
		Description: common.Qualify,
	})
	if err != nil {
		log.Errorf("GetKeywordRank: %v error: %v", keyword, err)
		return nil
	}

	resUnqualify, err := postgres.KeywordRankRepo().FindAll(ctx, &model.KeywordRankFindQuery{
		Keyword:     keyword,
		Description: common.UnQualify,
	})
	if err != nil {
		log.Errorf("GetKeywordRank: %v error: %v", keyword, err)
		return nil
	}

	lst := []model.GetKeywordRankResponse{}
	for _, v := range resQualify {
		lst = append(lst, model.GetKeywordRankResponse{
			Keyword: v.Keyword,
			Rank:    v.Rank,
			Url:     v.Url,
			Title:   v.Title,
		})
	}

	for _, v := range resUnqualify {
		lst = append(lst, model.GetKeywordRankResponse{
			Keyword: v.Keyword,
			Rank:    v.Rank,
			Url:     v.Url,
			Title:   v.Title,
		})
	}

	return lst
}
