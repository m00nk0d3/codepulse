package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/m00nk0d3/codePulse/internal/githubapi"
	"golang.org/x/oauth2"
)

// HandleGitHubLogin redireciona o user para o GitHub para login
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	OAuth2Config := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"repo", "notifications"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	url := OAuth2Config.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

// HandleGitHubCallback processes the response from login
func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}
	OAuth2Config := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"repo", "notifications"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	token, err := OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusInternalServerError)
		return
	}
	client := githubapi.NewClient(token.AccessToken)
	user, err := githubapi.GetUserInfo(client)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch user: %v", err), http.StatusInternalServerError)
		return
	}
	// shows the user name on GitHub
	fmt.Fprintf(w, "Hello, %s!", *user.Login)
}
