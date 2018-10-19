package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	
	"github.com/bwmarrin/discordgo"
	"github.com/andersgl/discordbot/conf"
	"github.com/andersgl/discordbot/bot/message"
	"github.com/andersgl/discordbot/bot/match"
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

func SendMessage(channel string, msg string) {
	DiscordConnection.ChannelMessageSend(channel, msg)
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
		msg := message.Parse(m)
		if len(msg.Command) > 0 {
			var response string

			if commandDisabled(msg.Command) >= 0 {
				return
			}
			
			switch msg.Command {
				case "enable":
					response = enableCommand(msg.Action)
				case "disable":
					response = disableCommand(msg.Action)
				case "roll":
					response = processCommand(Roll{}, msg)
				case "prac":
					response = processCommand(Prac{}, msg)
				case "match":
					response = processCommand(match.Match{}, msg)
			}

			s.ChannelMessageSend(m.ChannelID, response)
		}		
	}
}

func processCommand(cmd Command, msg message.Message) string {
	return cmd.Process(msg)
}

func commandDisabled(cmd string) int {
    for k, a := range Config.DisabledCmds {
        if a == cmd {
            return k
        }
    }
    return -1
}

func enableCommand(cmd string) string {
	index := commandDisabled(cmd)
	if index >= 0 {
		Config.DisabledCmds = append(Config.DisabledCmds[:index], Config.DisabledCmds[index+1:]...)
		return cmd + " is now enabled."
	}
	fmt.Println("enabled", Config.DisabledCmds)
	return cmd + " already enabled."
}

func disableCommand(cmd string) string {
	if commandDisabled(cmd) == -1 {
		Config.DisabledCmds = append(Config.DisabledCmds, cmd)
		return cmd + " is now disabled."
	}
	fmt.Println("disabled", Config.DisabledCmds)
	return cmd + " already disabled."
}

type Help struct {
	cmd string
	desc string
}

type Command interface {
    Process(msg message.Message) string
}