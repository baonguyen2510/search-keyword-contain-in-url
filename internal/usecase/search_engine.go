package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"search-keyword-service/common"
	"search-keyword-service/configs"
	"search-keyword-service/pkg/log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type SearchEngineServiceInterface interface {
	SearchKeyword(ctx context.Context, keyword string) (results []map[string]interface{}, err error)
}

var searchEngineService SearchEngineServiceInterface = &SearchEngineServiceImpl{}

type SearchEngineServiceImpl struct {
}

func SearchEngineService() SearchEngineServiceInterface {
	return searchEngineService
}

func (se SearchEngineServiceImpl) SearchKeyword(ctx context.Context, keyword string) (results []map[string]interface{}, err error) {

	c := colly.NewCollector(
		colly.AllowedDomains(
			fmt.Sprintf("www.%v.com", configs.Config.SearchEngineAddr),
			fmt.Sprintf("%v.com", configs.Config.SearchEngineAddr),
		),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	}

	// Set random User-Agent for each request
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
		log.Infof("Visiting: %v", r.URL.String())
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  fmt.Sprintf("*%v.*", configs.Config.SearchEngineAddr),
		Delay:       5 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Get the link from the href attribute
		link := e.Attr("href")

		// Extract the actual URL from the redirect link
		if strings.HasPrefix(link, "/url?q=") {
			parsedLink, err := url.Parse(link)
			if err == nil {
				// Get the query parameters
				queryParams := parsedLink.Query()
				// Get the actual URL
				link = queryParams.Get("q")
			}
		}

		// Get the title from the h3 tag
		title := e.ChildText("h3")

		parsedUrl, err := url.Parse(link)
		if err != nil {
			log.Errorf("Error parsing URL: err %v", err)
			return
		}

		isQualified := common.UnQualify
		if strings.Contains(strings.ToLower(parsedUrl.Host), strings.ToLower(keyword)) {
			isQualified = common.Qualify
		}

		// Append result to the slice with rank included
		results = append(results, map[string]interface{}{
			"title":     title,
			"url":       parsedUrl.String(),
			"qualified": isQualified,
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorf("Request URL: %v failed with response: %v, error: %v", r.Request.URL, string(r.Body), err)
	})

	// Simulate visiting search page
	searchUrl := fmt.Sprintf("https://www.%v.com/search?q=%s", configs.Config.SearchEngineAddr, keyword)
	err = c.Visit(searchUrl)
	if err != nil {
		log.Errorf("Error visiting search page: err %v", err)
		return nil, err
	}

	rankedResults := rankResults(results)

	return rankedResults, nil
}

func rankResults(results []map[string]interface{}) []map[string]interface{} {
	var (
		lstFirstValue  []map[string]interface{}
		lstSecondValue []map[string]interface{}
		lstThirdValue  []map[string]interface{}
		lstUnqualified []map[string]interface{}
	)

	for _, result := range results {
		if result["qualified"] == common.Qualify && result["title"] != "" {
			lstFirstValue = append(lstFirstValue, result)
		} else if result["qualified"] == common.Qualify && result["title"] == "" {
			lstSecondValue = append(lstSecondValue, result)
		} else if result["qualified"] == common.UnQualify && result["title"] != "" {
			lstThirdValue = append(lstThirdValue, result)
		} else if result["qualified"] == common.UnQualify {
			lstUnqualified = append(lstUnqualified, result)
		}
	}

	finalResults := make([]map[string]interface{}, 0, len(results))

	// Add all qualified results with sequential ranks starting from 1
	for i, v := range lstFirstValue {
		finalResults = append(finalResults, map[string]interface{}{
			"rank":      i + 1,
			"title":     v["title"],
			"url":       v["url"],
			"qualified": v["qualified"],
			"group":     1,
		})
	}

	rank := len(lstFirstValue) + 1
	for _, v := range lstSecondValue {
		finalResults = append(finalResults, map[string]interface{}{
			"rank":      rank,
			"title":     v["title"],
			"url":       v["url"],
			"qualified": v["qualified"],
			"group":     2,
		})
		rank++
	}

	rank = len(lstFirstValue) + len(lstSecondValue) + 1
	for _, v := range lstThirdValue {
		finalResults = append(finalResults, map[string]interface{}{
			"rank":      rank,
			"title":     v["title"],
			"url":       v["url"],
			"qualified": v["qualified"],
		})
		rank++
	}

	// Add unqualified results with incremented ranks starting from next after qualified
	rank = len(lstFirstValue) + len(lstSecondValue) + len(lstThirdValue) + 1
	for _, v := range lstUnqualified {
		finalResults = append(finalResults, map[string]interface{}{
			"rank":      rank,
			"title":     v["title"],
			"url":       v["url"],
			"qualified": v["qualified"],
		})
		rank++
	}

	return finalResults
}
