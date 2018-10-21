package forlulz

import (
	"github.com/andersgl/discordbot/message"
)

func LOL(msg *message.Message) {
	if msg.Content() == "cawer" {
		msg.Respond("KING CAWER!!!")
	}
}