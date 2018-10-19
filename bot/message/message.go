package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type User struct {
	Id string
	Username string
}

type Message struct {
	Command string
	Action string
	Args []string
	Channel string
	User User
}

func Parse(m *discordgo.MessageCreate) Message {
	user := User{Id: m.Author.ID, Username: m.Author.Username}
	msg := Message{User: user, Channel: m.ChannelID}
	parts := strings.Split(m.Content[1:], " ")
	if len(parts) > 0 {
		msg.Command = parts[0]
	}
	if len(parts) > 1 {
		msg.Action = parts[1]
	}
	if len(parts) > 2 {
		msg.Args = parts[2:]
	}
	return msg
}