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
        ctx.HTMLSet(200, "base", "index")
    })
    m.Get("/bots/:id", func(ctx *macaron.Context) {
        ctx.HTMLSet(200, "base", "bot")
    })
    err = http.ListenAndServe(conf.http.port, m)
    if err != nil {
        log.Fatal(err)
    }
}
