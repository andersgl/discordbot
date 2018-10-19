package main

import (
	"flag"

	"github.com/andersgl/discordbot/conf"
	"github.com/andersgl/discordbot/bot"
)

// Variables used for command line parameters
var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	config, _ := conf.Load()
	if len(token) > 0 {
		config.Token = token
	}
	bot.Start(config)
}
