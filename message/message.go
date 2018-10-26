package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type User struct {
	Id string
	Username string
	Admin bool
}

type Message struct {
	IsCommand bool
	Command string
	Action string
	Args []string
	User *User
	MessageCreate *discordgo.MessageCreate
	Session *discordgo.Session
}

func (msg Message) Respond(response string) {
	msg.Session.ChannelMessageSend(msg.MessageCreate.ChannelID, response)
}

func (msg Message) RespondTTS(response string) {
	msg.Session.ChannelMessageSendTTS(msg.MessageCreate.ChannelID, response)
}

func (msg Message) Content() string {
	return msg.MessageCreate.Content
}

func New(m *discordgo.MessageCreate, s *discordgo.Session, trigger string, admins []string) Message {
	user := User{
		Id: m.Author.ID, 
		Username: m.Author.Username,
		Admin: userIsAdmin(m.Author.ID, admins)}
	msg := Message{
		IsCommand: strings.HasPrefix(m.Content, trigger), 
		User: &user, 
		MessageCreate: m,
		Session: s}
	parts := strings.Split(m.Content[1:], " ")
	if len(parts) > 0 {
		msg.Command = strings.ToLower(parts[0])
	}
	if len(parts) > 1 {
		msg.Action = strings.ToLower(parts[1])
	}
	if len(parts) > 2 {
		msg.Args = parts[2:]
	}
	return msg
}

func userIsAdmin(userId string, admins []string) bool {
	for _, id := range admins {
        if id == userId {
            return true
        }
    }
	return false
}