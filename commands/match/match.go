package match

import (
	"github.com/andersgl/discordbot/message"
	"github.com/andersgl/discordbot/help"
)

type Match struct {}

func (m Match) Process(msg *message.Message) {
	switch msg.Action {
		case "help":
			m.help(msg)
		case "yes", "no":
			m.update(msg)
		default:
			m.summary(msg)
	}
}

func (m Match) summary(msg *message.Message) {
	msg.Respond("Summary")
}

func (m Match) update(msg *message.Message) {
	msg.Respond("Summary")
}

func (m Match) help(msg *message.Message) {
	response := "**Commands:**\n"
	helpers := []help.Help{
		{"!match", "get a prac summary"},
		{"!match yes <game>", "sign up for prac today (game is optional)"},
		{"!match no <game>", "let people know you can't prac today (game is optional)"},
		{"!match remove <game>", "remove yourself from the prac entry (game is optional)"},
	}
	for _, helper := range helpers {
		response += "**" + helper.Cmd + "** - " + helper.Desc + "\n"
    }
	msg.Respond(response)
}

func New() Match {
	return Match{}
}