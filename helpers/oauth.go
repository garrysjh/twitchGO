package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Oauth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func OauthGenerator() string {
	resp, err := http.PostForm("https://id.twitch.tv/oauth2/token", url.Values{
		"client_id":     {EnvVariable("CLIENT_ID")},
		"client_secret": {EnvVariable("CLIENT_SECRET")},
		"grant_type":    {"client_credentials"},
	})
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result Oauth
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result.AccessToken
}
