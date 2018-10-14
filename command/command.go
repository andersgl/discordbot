package command

import (
	"github.com/bwmarrin/discordgo"
)

var Msg *discordgo.MessageCreate

func SetMessage(m *discordgo.MessageCreate) {
	Msg = m
}

func Run() string {
	if Msg.Content == "!roll" {
		return Roll()
	}
	if Msg.Content == "ping" {
		return Ping()
	}

	return ""
}
