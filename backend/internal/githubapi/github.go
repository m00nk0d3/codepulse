package githubapi

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client := github.NewClient(oauth2.NewClient(context.Background(), ts))
	return client
}

func GetUserInfo(client *github.Client) (*github.User, error) {
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}
	return user, nil
}
