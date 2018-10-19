package bot

import (
	// "fmt"
)

type Prac struct {}

func (p Prac) Process(msg Message) string {
	switch msg.action {
		case "help":
			return p.help()
		case "yes", "no":
			return p.update(msg)
		default:
			return p.summary(msg)
	}
}

func (p Prac) summary(msg Message) string {
	return "Summary"
}

func (p Prac) update(msg Message) string {
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
