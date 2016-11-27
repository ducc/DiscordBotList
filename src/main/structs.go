package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type (
	botModel struct {
		id          string
		name        string
		description string
		category    string
		website     string
		server      string
		authorize   string
		avatar      string
		online      bool
		ranking     int
		servers     int
		featured    bool
		score       int
		reviews     int
	}
	config struct {
		Http struct {
			Port string `json:"port"`
		} `json:"http"`
		Templates struct {
			Directory string `json:"directory"`
		} `json:"templates"`
		OAuth struct {
			ClientId         string   `json:"client_id"`
			ClientSecret     string   `json:"client_secret"`
			Scopes           []string `json:"scopes"`
			RedirectUrl      string   `json:"redirect_url"`
			AuthorizationUrl string   `json:"authorization_url"`
			TokenUrl         string   `json:"token_url"`
		} `json:"oauth"`
	}
	oauthClient struct {
		tokenResponse oauthTokenResponse
		userResponse  *oauthUserResponse
	}
	oauthGuildResponse struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Icon        string `json:"icon"`
		Owner       bool   `json:"owner"`
		Permissions int    `json:"permissions"`
	}
	oauthTokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}
	oauthUserResponse struct {
		Id            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"descriminator"`
		Avatar        string `json:"avatar"`
		Bot           bool   `json:"bot"`
		MfaEnabled    bool   `json:"mfa_enabled"`
	}
	user struct {
		client *oauthClient
	}
)

func newOAuthClient(tokenResponse oauthTokenResponse) *oauthClient {
	c := new(oauthClient)
	c.tokenResponse = tokenResponse
	return c
}

func (client *oauthClient) refresh() {
	// TODO new access token using refresh token
}

func (client *oauthClient) request(method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Add("Authorization", "Bearer "+client.tokenResponse.AccessToken)
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}

func (client *oauthClient) user(useCache bool) *oauthUserResponse {
	if useCache && client.userResponse != nil {
		return client.userResponse
	}
	resp := client.request("GET", discord_api_user_url, nil)
	var userResponse oauthUserResponse
	temp := &userResponse
	json.NewDecoder(resp.Body).Decode(temp)
	client.userResponse = temp
	return temp
}

func (client *oauthClient) guilds() []oauthGuildResponse {
	resp := client.request("GET", discord_api_guilds_url, nil)
	guildResponses := make([]oauthGuildResponse, 0)
	json.NewDecoder(resp.Body).Decode(&guildResponses)
	return guildResponses
}

func newUser(client *oauthClient) *user {
	u := new(user)
	u.client = client
	return u
}
