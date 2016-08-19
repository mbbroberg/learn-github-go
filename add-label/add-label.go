package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

const keyword = "snap-plugin"

var (
	personalAccessToken string
	// issuesCollection    allIssues
	org string
)

// TokenSource is an encapsulation of the AccessToken string
type TokenSource struct {
	AccessToken string
}

// Token authenticates via oauth
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func GetRepos() (allRepos []*github.Repository) {
	personalAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")

	if len(personalAccessToken) == 0 {
		log.Fatal("Before you can use this you must set the GITHUB_ACCESS_TOKEN environment variable.")
	}

	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient) // authenticated to GitHub here

	org := "intelsdi-x"
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	for {
		repos, resp, err := client.Repositories.ListByOrg(org, opt)
		if err != nil {
			fmt.Println(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
	return allRepos
}

func LabelRepos(allRepos []*github.Repository) {
	personalAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")

	if len(personalAccessToken) == 0 {
		log.Fatal("Before you can use this you must set the GITHUB_ACCESS_TOKEN environment variable.")
	}

	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient) // authenticated to GitHub here
	color := "ededed"
	labelName := "tracked"
	for _, rp := range allRepos {
		repo := *rp.Name
		if strings.Contains(repo, "snap-plugin-collector-use") {
			label, _, err := client.Issues.GetLabel("intelsdi-x", repo, "tracked")
			if err != nil {
				fmt.Printf("Tracked label already exists for %v", repo)
				break
			}
			fmt.Printf("Repo: %v, Label %v, error: %v\n", repo, label, err)
			if *label.Color != "ededed" {
				_, _, err := client.Issues.EditLabel("intelsdi-x", repo, labelName, &github.Label{Name: &labelName, Color: &color})
				if err != nil {
					fmt.Println("Wooot.")
				}
			}
		}
	}
}

func main() {
	repos := GetRepos()
	LabelRepos(repos)
}
