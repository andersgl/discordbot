package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	
	"github.com/bwmarrin/discordgo"
	"github.com/andersgl/discordbot/conf"
)

var (
	DiscordConnection *discordgo.Session
	Config conf.Conf
)

func Start(config conf.Conf) {
	Config = config

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Config.Token)
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

	if strings.HasPrefix(m.Content, Config.CommandTrigger) {
		msg := parseMessage(m)
		if len(msg.command) > 0 {
			var response string

			if CommandDisabled(msg.command) >= 0 {
				return
			}
			
			switch msg.command {
				case "enable":
					response = EnableCommand(msg.action)
				case "disable":
					response = DisableCommand(msg.action)
				case "roll":
					response = processCommand(Roll{}, msg)
				case "prac":
					response = processCommand(Prac{}, msg)
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

func processCommand(cmd Command, msg Message) string {
	return cmd.Process(msg)
}

func CommandDisabled(cmd string) int {
    for k, a := range Config.DisabledCmds {
        if a == cmd {
            return k
        }
    }
    return -1
}

func EnableCommand(cmd string) string {
	index := CommandDisabled(cmd)
	if index >= 0 {
		Config.DisabledCmds = append(Config.DisabledCmds[:index], Config.DisabledCmds[index+1:]...)
		return cmd + " is now enabled."
	}
	fmt.Println("enabled", Config.DisabledCmds)
	return cmd + " already enabled."
}

func DisableCommand(cmd string) string {
	if CommandDisabled(cmd) == -1 {
		Config.DisabledCmds = append(Config.DisabledCmds, cmd)
		return cmd + " is now disabled."
	}
	fmt.Println("disabled", Config.DisabledCmds)
	return cmd + " already disabled."
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

type Help struct {
	cmd string
	desc string
}

type Command interface {
    Process(msg Message) string
}