package roll

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/andersgl/discordbot/help"
	"github.com/andersgl/discordbot/message"
)

type Roll struct{}

type rollResult struct {
	user   *message.User
	result int
}

var rollResults = make(map[string]rollResult)
var isRolling bool = false

func (r Roll) Process(msg *message.Message) {
	switch msg.Action {
	case "start":
		r.startContest(msg)
	case "help":
		r.help(msg)
	default:
		r.rollNow(msg)
	}
}

func (r Roll) rollNow(msg *message.Message) {
	result := randomInt(1, 100)
	if isRolling {
		_, hasRolled := rollResults[msg.User.Id]
		if hasRolled {
			msg.Respond("You can only roll once!")
		} else {
			rollResults[msg.User.Id] = rollResult{msg.User, result}
			msg.Respond(msg.User.Username + " rolled, wait for results.")
		}
	} else {
		msg.Respond(msg.User.Username + " rolls ... " + strconv.Itoa(result))
	}
}

func (r Roll) startContest(msg *message.Message) {
	if isRolling {
		msg.Respond("A contest is already rollin' ...")
	}
	timeout := 10
	if len(msg.Args) > 0 {
		userTimeout, _ := strconv.Atoi(msg.Args[0])
		if userTimeout > 0 {
			timeout = userTimeout
		}
		if timeout > 60 {
			timeout = 60
		}
	}
	isRolling = true
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		r.finishContest(msg)
	})

	var prize string
	if len(msg.Args) > 1 {
		prize = strings.Join(msg.Args[1:], " ")
	}

	response := "Starting !roll contest."
	if len(prize) > 0 {
		response += " Prize: " + prize + "."
	}
	response += "\nResults will be announced in " + strconv.Itoa(timeout) + " seconds. Start rollin' rollin' rollin' ..."
	msg.Respond(response)
}

func (r Roll) finishContest(msg *message.Message) {
	isRolling = false
	response := "``` Results\n"
	var max rollResult
	for _, value := range rollResults {
		response += value.user.Username + ": " + strconv.Itoa(value.result) + "\n"
		if value.result > max.result {
			max = value
		}
	}
	response += "... and the winner is: " + max.user.Username
	response += "```"
	rollResults = make(map[string]rollResult)

	msg.Respond(response)
}

func (r Roll) help(msg *message.Message) {
	response := "**Commands:**\n"
	helpers := []help.Help{
		{"!roll", "just roll"},
		{"!roll start <time> <prize>", "start a roll contest for a specified time and prize"},
	}
	for _, helper := range helpers {
		response += "**" + helper.Cmd + "** - " + helper.Desc + "\n"
	}
	msg.Respond(response)
}

func New() Roll {
	return Roll{}
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
