package main

import (
	"flag"

	"github.com/andersgl/discordbot/bot"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bot.Connect(Token)
}
