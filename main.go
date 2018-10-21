package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/andersgl/discordbot/config"
	"github.com/andersgl/discordbot/message"
	"github.com/andersgl/discordbot/forlulz"

	// Command packages go here
	"github.com/andersgl/discordbot/commands/roll"
	"github.com/andersgl/discordbot/commands/prac"
	"github.com/andersgl/discordbot/commands/match"
)

// Variables used for command line parameters
var (
	token string
	conf config.Config
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	conf = config.Load()
	if len(token) > 0 {
		conf.Token = token
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	var response string
	msg := message.New(m, s, conf.CommandTrigger, conf.Admins)
	if msg.IsCommand {
		if commandDisabled(msg.Command) >= 0 {
			return
		}
		
		switch msg.Command {
			// Enable a command
			case "enable":
				response = enableCommand(msg.Action)
			// Disable a command
			case "disable":
				response = disableCommand(msg.Action)
			
			// Specific commands
			case "roll":
				response = processCommand(roll.New(), &msg)
			case "prac":
				response = processCommand(prac.New(), &msg)
			case "match":
				response = processCommand(match.New(), &msg)
		}		
	} else {
		response = forlulz.LOL(&msg)
	}

	if len(response) > 0 {
		s.ChannelMessageSend(m.ChannelID, response)
	}
}

func processCommand(cmd Command, msg *message.Message) string {
	return cmd.Process(msg)
}

func commandDisabled(cmd string) int {
    for k, a := range conf.DisabledCmds {
        if a == cmd {
            return k
        }
    }
    return -1
}

func enableCommand(cmd string) string {
	index := commandDisabled(cmd)
	if index >= 0 {
		conf.DisabledCmds = append(conf.DisabledCmds[:index], conf.DisabledCmds[index+1:]...)
		return cmd + " is now enabled."
	}
	return cmd + " already enabled."
}

func disableCommand(cmd string) string {
	if commandDisabled(cmd) == -1 {
		conf.DisabledCmds = append(conf.DisabledCmds, cmd)
		return cmd + " is now disabled."
	}
	return cmd + " already disabled."
}

type Command interface {
    Process(msg *message.Message) string
}
