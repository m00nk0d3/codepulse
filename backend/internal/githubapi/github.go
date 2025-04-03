package githubapi

import (
	"context" // Adicionando o contexto

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewClient cria um cliente autenticado para a API do GitHub
func NewClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client := github.NewClient(oauth2.NewClient(context.Background(), ts)) // Corrigido com context.Background()
	return client
}

// GetUserInfo retorna as informações do usuário autenticado
func GetUserInfo(client *github.Client) (*github.User, error) {
	user, _, err := client.Users.Get(context.Background(), "") // Corrigido com context.Background()
	if err != nil {
		return nil, err
	}
	return user, nil
}
