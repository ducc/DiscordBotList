package main

import (
    "log"
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/pongo2"
    "net/http"
    "io/ioutil"
)

const (
    config_path = "config.json"
)

var (
    conf *config
)

func main() {
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
    m.Get("/", func(ctx *macaron.Context) {
        ctx.Data["featured_bots"] = []botModel{
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
        }
        ctx.Data["top_rated_bots"] = []botModel{
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
        }
        ctx.Data["recently_added_bots"] = []botModel{
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
            {id: "123456789", name: "Vonodosh Bot", category: "Games", description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.", website: "http://bit.ly/vonodosh", server: "PVHQxBd", authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519", avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg", online: true, ranking: 53, servers: 484, featured: false, score: 4, reviews: 835, },
        }
        ctx.HTMLSet(200, "base", "index")
    })
    m.Get("/bots/:id", func(ctx *macaron.Context) {
        ctx.Data["bot"] = botModel{
            id: "123456789",
            name: "Vonodosh Bot",
            category: "Games",
            description: "Vonodosh. A bot that can fulfill your entertaining needs, with games to keep your guild busy.",
            website: "http://bit.ly/vonodosh",
            server: "PVHQxBd",
            authorize: "https://discordapp.com/oauth2/authorize?client_id=192332191240421377&scope=bot&permissions=536083519",
            avatar: "https://cdn.discordapp.com/app-icons/170242612425392128/8744452b039eaec1ba34c0a96c3bf428.jpg",
            online: true,
            ranking: 53,
            servers: 484,
            featured: false,
            score: 4,
            reviews: 835,
        }
        ctx.Data["content"] = "<img src=\"http://i.imgur.com/lvfXcZm.png\" class=\"image\">"
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
        body, _ := ioutil.ReadAll(resp.Body)
        log.Println(string(body))
    })
    err = http.ListenAndServe(conf.Http.Port, m)
    if err != nil {
        log.Fatal(err)
    }
}
