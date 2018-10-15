package bot

import (
	"math/rand"
	"strconv"
	"time"
	"strings"
)

type Roll struct {}

type rollResult struct {
	user User
	result int
}

var rollResults = make(map[string]rollResult)
var isRolling bool = false

func (r Roll) Process(msg Message) string {
	switch msg.action {
		case "start":
			return r.startContest(msg)
		default:
			return r.rollNow(msg.user)
	}
}

func (r Roll) rollNow(user User) string {
	result := randomInt(1,100)
	if isRolling {
		_, ok := rollResults[user.id]
		if ok {
			return "You can only roll once!"
		} else {
			rollResults[user.id] = rollResult{user, result}
			return user.username + " rolled, wait for results."
		}
	}
	return user.username + " rolls ... " + strconv.Itoa(result)
}

func (r Roll) startContest(msg Message) string {
	if isRolling {
		return "A contest is already rollin' ..."
	}
	timeout := 10
	if len(msg.args) > 0 {
		userTimeout, _ := strconv.Atoi(msg.args[0])
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
	if len(msg.args) > 1 {
		prize = strings.Join(msg.args[1:], " ")
	}

	response := "Starting !roll contest."
	if len(prize) > 0 {
		response += " Prize: " + prize + "."
	}
	response += "\nResults will be announced in " + strconv.Itoa(timeout) + " seconds. Start rollin' rollin' rollin' ..."
	return response
}

func (r Roll) finishContest(msg Message) {
	isRolling = false
	response := "``` Results\n"
	var max rollResult
	for _, value := range rollResults {
		response += value.user.username + ": " + strconv.Itoa(value.result) + "\n"
		if value.result > max.result {
			max = value
		}
	}
	response += "... and the winner is: " + max.user.username
	response += "```"
	rollResults = make(map[string]rollResult)
	SendMessage(msg.channel, response)
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}