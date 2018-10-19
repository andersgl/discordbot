package bot

import (
	// "fmt"
	"github.com/andersgl/discordbot/bot/message"
)

type Prac struct {}

func (p Prac) Process(msg message.Message) string {
	switch msg.Action {
		case "help":
			return p.help()
		case "yes", "no":
			return p.update(msg)
		default:
			return p.summary(msg)
	}
}

func (p Prac) summary(msg message.Message) string {
	return "Summary"
}

func (p Prac) update(msg message.Message) string {
	return "Summary"
}

func (p Prac) help() string {
	response := "**Commands:**\n"
	helpers := []Help{
		{"!prac", "get a prac summary"},
		{"!prac yes <game>", "sign up for prac today (game is optional)"},
		{"!prac no <game>", "let people know you can't prac today (game is optional)"},
		{"!prac remove <game>", "remove yourself from the prac entry (game is optional)"},
	}
	for _, helper := range helpers {
		response += "**" + helper.cmd + "** - " + helper.desc + "\n"
    }
	return response
}
