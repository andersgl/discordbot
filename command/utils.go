package command

import (
	"math/rand"
	"strconv"
	"time"
)

func intToString(value int) string {
	return strconv.Itoa(value)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}