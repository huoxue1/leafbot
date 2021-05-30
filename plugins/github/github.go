package github

import (
	"fmt" //nolint:gci
	"github.com/google/go-github/v35/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	client *github.Client
)

var tokenContent string

func Init(token string) {
	if token != "" {
		tokenContent = token
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tokenContent},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
	responses, err := SearchResponse(ctx)
	if err != nil {
		log.Errorln("获取信息失败")
		return
	}
	for _, respons := range responses {
		log.Infoln(respons)
	}
}

func SearchResponse(ctx context.Context) ([]string, error) {
	var msgs []string
	repositories, _, err := client.Search.Repositories(ctx, "LeafBot", &github.SearchOptions{})
	if repositories == nil {
		return nil, err
	}
	for _, repository := range repositories.Repositories {

		msg := fmt.Sprintf("%v\nDescription: %v\nStar/Fork/Issue: %d / %d / %d\nLanguage: %v\nLicense: \nLastPushed: %v\nJump:%v",
			repository.GetName(),
			repository.GetDescription(),
			repository.GetStargazersCount(),
			repository.GetForksCount(),
			repository.GetOpenIssuesCount(),
			repository.GetLanguage(),
			//repository.GetLicense().String(),
			repository.GetPushedAt().Format("2006-01-02 15:04:05"),
			repository.GetURL(),
		)
		msgs = append(msgs, msg)
	}

	return msgs, err
}
