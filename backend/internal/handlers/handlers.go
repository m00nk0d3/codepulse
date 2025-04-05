package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/m00nk0d3/codePulse/internal/githubapi"
	"github.com/m00nk0d3/codePulse/internal/helpers"
)

// HandleGitHubLogin redirects the user to GitHub for login
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	authURL := helpers.OAuth2Config.AuthCodeURL("")
	http.Redirect(w, r, authURL, http.StatusFound)
}

// exchangeToken exchanges the authorization code for an access token
func exchangeToken(code string) (*oauth2.Token, error) {
	token, err := helpers.OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging token: %v", err)
		return nil, err
	}
	return token, nil
}

// fetchUserInfo fetches the user info from GitHub
func fetchUserInfo(token *oauth2.Token) (*github.User, error) {
	client := githubapi.NewClient(token.AccessToken)
	user, err := githubapi.GetUserInfo(client)
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		return nil, err
	}
	return user, nil
}

// HandleGitHubCallback processes the response from login
func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")
	if authCode == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	token, err := exchangeToken(authCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusInternalServerError)
		return
	}

	user, err := fetchUserInfo(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch user: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", *user.Login)
}
