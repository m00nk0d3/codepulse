package auth

import (
	"context"
	"log"
	"os"

	// Adicionando o contexto
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// OAuth2Config guarda a configuração do OAuth
var OAuth2Config = oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"repo", "notifications"},
	Endpoint:     github.Endpoint,
}

// GetOAuth2Token troca o código pela chave de acesso
func GetOAuth2Token(code string) (*oauth2.Token, error) {
	token, err := OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Failed to exchange token:", err)
		return nil, err
	}
	return token, nil
}
