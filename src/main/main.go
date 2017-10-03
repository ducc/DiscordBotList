package main

import (
	"encoding/json"
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
	"time"
	"fmt"
)

const (
	config_path = "config.json"
)

var (
	conf  *config
	users = make(map[string]*user)
)

func main() {
	
	fmt.Println( "Sort your life out, use dank bot list." )
	time.Sleep( time.Second * 2 )
	
	var err error
	conf, err = loadConfig(config_path)
	if err != nil {
		log.Fatal(err)
	}
	m := macaron.Classic()
	m.Use(macaron.Static("static", macaron.StaticOptions{
		Prefix: "static",
	}))
	m.Use(pongo2.Pongoers(pongo2.Options{
		Directory: conf.Templates.Directory,
	}, "base:templates"))
	demoBotModel := botModel{id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835}
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["featured_bots"] = []botModel{
			demoBotModel, demoBotModel, demoBotModel, demoBotModel, demoBotModel,
		}
		ctx.Data["top_rated_bots"] = []botModel{
			demoBotModel, demoBotModel, demoBotModel, demoBotModel, demoBotModel,
		}
		ctx.Data["recently_added_bots"] = []botModel{
			demoBotModel, demoBotModel, demoBotModel, demoBotModel, demoBotModel,
		}
		ctx.HTMLSet(200, "base", "index")
	})
	m.Get("/bots/:id", func(ctx *macaron.Context) {
		ctx.Data["bot"] = demoBotModel
		ctx.Data["content"] = "<div class=\"box\"><p>This is some example content!</p></div>"
		ctx.HTMLSet(200, "base", "bot")
	})
	m.Get("/login", func(ctx *macaron.Context) {
		ctx.Redirect(getOAuthAuthorizationUrl(), 302)
	})
	m.Get("/redirect", func(ctx *macaron.Context) {
		code := ctx.Query("code")
		req, err := http.NewRequest("POST", getOAuthTokenUrl(code), nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		var tokenResponse oauthTokenResponse
		json.NewDecoder(resp.Body).Decode(&tokenResponse)
		u := newUser(newOAuthClient(tokenResponse))
		users[u.client.user(true).Id] = u
	})
	err = http.ListenAndServe(conf.Http.Port, m)
	if err != nil {
		log.Fatal(err)
	}
}
