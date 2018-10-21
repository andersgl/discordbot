package prac

import (
	"github.com/andersgl/discordbot/message"
	"github.com/andersgl/discordbot/help"
)

type Prac struct {}

func (p Prac) Process(msg *message.Message) {
	switch msg.Action {
		case "help":
			p.help(msg)
		case "yes", "no":
			p.update(msg)
		default:
			p.summary(msg)
	}
}

func (p Prac) summary(msg *message.Message) {
	msg.Respond("Summary")
}

func (p Prac) update(msg *message.Message) {
	msg.Respond("Summary")
}

func (p Prac) help(msg *message.Message) {
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
	msg.Respond(response)
}

func New() Prac {
	return Prac{}
}
