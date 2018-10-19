package bot

import (
	"math/rand"
	"strconv"
	"time"
	"strings"
	"github.com/andersgl/discordbot/bot/message"
)

type Roll struct {}

type rollResult struct {
	user message.User
	result int
}

var rollResults = make(map[string]rollResult)
var isRolling bool = false

func (r Roll) Process(msg message.Message) string {
	switch msg.Action {
		case "start":
			return r.startContest(msg)
		case "help":
			return r.help()
		default:
			return r.rollNow(msg.User)
	}
}

func (r Roll) rollNow(user message.User) string {
	result := randomInt(1,100)
	if isRolling {
		_, ok := rollResults[user.Id]
		if ok {
			return "You can only roll once!"
		} else {
			rollResults[user.Id] = rollResult{user, result}
			return user.Username + " rolled, wait for results."
		}
	}
	return user.Username + " rolls ... " + strconv.Itoa(result)
}

func (r Roll) startContest(msg message.Message) string {
	if isRolling {
		return "A contest is already rollin' ..."
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
	return response
}

func (r Roll) finishContest(msg message.Message) {
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
	SendMessage(msg.Channel, response)
}

func (r Roll) help() string {
	response := "**Commands:**\n"
	helpers := []Help{
		{"!roll", "just roll"},
		{"!roll start <time> <prize here>", "start a roll contest for a specified time and prize"},
	}
	for _, helper := range helpers {
		response += "**" + helper.cmd + "** - " + helper.desc + "\n"
    }
	return response
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}