package forlulz

import (
	"github.com/andersgl/discordbot/message"
)

func LOL(msg *message.Message) string {
	if msg.Content() == "cawer" {
		return "KING CAWER!!!"
	}
	
	return ""
}