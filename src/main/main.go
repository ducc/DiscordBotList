package main

import (
    "log"
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/pongo2"
    "net/http"
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
        Directory: conf.templates.directory,
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
    err = http.ListenAndServe(conf.http.port, m)
    if err != nil {
        log.Fatal(err)
    }
}
