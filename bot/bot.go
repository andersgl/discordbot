package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	
	"github.com/bwmarrin/discordgo"
)

var DiscordConnection *discordgo.Session

func Connect(token string) {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func SendMessage(channel string, message string) {
	DiscordConnection.ChannelMessageSend(channel, message)
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	DiscordConnection = s
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		msg := parseMessage(m)
		if len(msg.command) > 0 {
			var response string
			
			switch msg.command {
				case "roll":
					response = processCommand(Roll{}, msg)
			}

			s.ChannelMessageSend(m.ChannelID, response)
		}		
	}
}

func parseMessage(m *discordgo.MessageCreate) Message {
	user := User{id: m.Author.ID, username: m.Author.Username}
	msg := Message{user: user, channel: m.ChannelID}
	parts := strings.Split(m.Content[1:], " ")
	if len(parts) > 0 {
		msg.command = parts[0]
	}
	if len(parts) > 1 {
		msg.action = parts[1]
	}
	if len(parts) > 2 {
		msg.args = parts[2:]
	}
	return msg
}

type User struct {
	id string
	username string
}

type Message struct {
	command string
	action string
	args []string
	channel string
	user User
}

type Command interface {
    Process(msg Message) string
}

func processCommand(cmd Command, msg Message) string {
	return cmd.Process(msg)
}