package auth

import (
	"context"
	"log"

	"github.com/m00nk0d3/codePulse/internal/helpers"
	"golang.org/x/oauth2"
)

// GetOAuth2Token troca o código pela chave de acesso
func GetOAuth2Token(code string) (*oauth2.Token, error) {
	token, err := helpers.OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Failed to exchange token:", err)
		return nil, err
	}
	return token, nil
}
