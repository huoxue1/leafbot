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
	responses := SearchResponse(ctx)
	for _, respons := range responses {
		log.Infoln(respons)
	}
}

func SearchResponse(ctx context.Context) []string {
	var msgs []string
	repositories, _, err := client.Search.Repositories(ctx, "LeafBot", &github.SearchOptions{})
	if err != nil {

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

	return msgs
}
