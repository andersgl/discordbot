package match

import (
	// "fmt"
	"github.com/andersgl/discordbot/bot/message"
)

type Match struct {}

func (m Match) Process(msg message.Message) string {
	switch msg.Action {
		case "help":
			return m.help()
		case "yes", "no":
			return m.update(msg)
		default:
			return m.summary(msg)
	}
}

func (m Match) summary(msg message.Message) string {
	return "Summary"
}

func (m Match) update(msg message.Message) string {
	return "Summary"
}

func (m Match) help() string {
	return "help"
	// response := "**Commands:**\n"
	// helpers := []bot.Help{
	// 	{"!match", "get a prac summary"},
	// 	{"!match yes <game>", "sign up for prac today (game is optional)"},
	// 	{"!match no <game>", "let people know you can't prac today (game is optional)"},
	// 	{"!match remove <game>", "remove yourself from the prac entry (game is optional)"},
	// }
	// for _, helper := range helpers {
	// 	response += "**" + helper.cmd + "** - " + helper.desc + "\n"
    // }
	// return response
}
