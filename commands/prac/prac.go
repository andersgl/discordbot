package prac

import (
	"github.com/andersgl/discordbot/message"
	"github.com/andersgl/discordbot/help"
)

type Prac struct {}

func (p Prac) Process(msg *message.Message) string {
	switch msg.Action {
		case "help":
			return p.help()
		case "yes", "no":
			return p.update(msg)
		default:
			return p.summary(msg)
	}
}

func (p Prac) summary(msg *message.Message) string {
	return "Summary"
}

func (p Prac) update(msg *message.Message) string {
	return "Summary"
}

func (p Prac) help() string {
	response := "**Commands:**\n"
	helpers := []help.Help{
		{"!prac", "get a prac summary"},
		{"!prac yes <game>", "sign up for prac today (game is optional)"},
		{"!prac no <game>", "let people know you can't prac today (game is optional)"},
		{"!prac remove <game>", "remove yourself from the prac entry (game is optional)"},
	}
	for _, helper := range helpers {
		response += "**" + helper.Cmd + "** - " + helper.Desc + "\n"
    }
	return response
}

func New() Prac {
	return Prac{}
}
