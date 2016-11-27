package main

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
)
