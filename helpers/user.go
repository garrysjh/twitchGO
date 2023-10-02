package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type baseData struct {
	Data []User `json:"data"`
}

type User struct {
	UserId       string `json:"id"`
	UserName     string `json:"login"`
	DisplayName  string `json:"display_name"`
	ProfilePic   string `json:"profile_image_url"`
	Description  string `json:"description"`
	ViewCount    int    `json:"view_count"`
	OfflineImage string `json:"offline_image_url"`
	CreatedAt    string `json:"created_at"`
}

func GetUserInfo(username string) User {
	client := http.Client{}
	var url string = "https://api.twitch.tv/helix/users?login=" + username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + OauthGenerator()},
		"Client-Id":     []string{EnvVariable("CLIENT_ID")},
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var userData baseData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		fmt.Println(err)
	}
	return userData.Data[0]
}

func GetUserDp(username string) string {
	return GetUserInfo(username).ProfilePic
}

func GetLoginName(username string) string {
	return GetUserInfo(username).UserName
}

func GetUserCreatedAt(username string) string {
	return GetUserInfo(username).CreatedAt
}
