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
		http struct {
			port string
		}
		templates struct {
			directory string
		}
		oauth struct {
			client_id         string
			client_secret     string
			scopes            []string
			redirect_url      string
			authorization_url string
			token_url         string
		}
	}
)
