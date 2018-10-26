package prac

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"time"

	"github.com/andersgl/discordbot/message"
	"github.com/andersgl/discordbot/help"
)

type Prac struct {
	games []string
}

type pracEntries struct {
	Yes []map[string]string `json:"yes"`
	No []map[string]string `json:"no"`
	Maybe []map[string]string `json:"maybe"`
}

type pracDate map[string]pracEntries

type jsonContent map[string]pracDate

const filePath string = "commands/prac/prac.json"

func (p Prac) Process(msg *message.Message) {
	switch msg.Action {
		case "help":
			p.help(msg)
		case "yes", "no":
			p.update(msg)
		default:
			p.summary(msg)
	}
}

func (p Prac) summary(msg *message.Message) {
	response := ""
	pracJson := p.loadJson()
	
	currentDate := time.Now().Local().Format("2006-01-02")
	
	if p.dateExists(currentDate, pracJson) {
		response = "Prac overview:\n"
		pracData := p.findDate(currentDate, pracJson)
		for game, pracEntries := range pracData {
			response += "```diff\n"
			response += strings.ToUpper(game) + "\n"

			var pracYes []string
			for _, entry := range pracEntries.Yes {
				for _, player := range entry {
					pracYes = append(pracYes, player)
				}
			}
			response += "+ " + strings.Join(pracYes[:], ", ") + "\n"

			var pracNo []string
			for _, entry := range pracEntries.No {
				for _, player := range entry {
					pracNo = append(pracNo, player)
				}
			}
			response += "- " + strings.Join(pracNo[:], ", ") + "\n"

			var pracMaybe []string
			for _, entry := range pracEntries.Maybe {
				for _, player := range entry {
					pracMaybe = append(pracMaybe, player)
				}
			}
			if len(pracMaybe) > 0 {
				response += "--- maybe: " + strings.Join(pracNo[:], ", ") + "\n"
			}

			response += "```"
		}
	} else {
		response = "No prac entries today :("
	}
	
	msg.Respond(response)
}

func (p Prac) update(msg *message.Message) {
	msg.Respond("Should update")
}

func (p Prac) help(msg *message.Message) {
	response := "**Commands:**\n"
	helpers := []help.Help{
		{"!prac", "get a prac summary"},
		{"!prac yes <game>", "sign up for prac today (game is optional)"},
		{"!prac no <game>", "let people know you can't prac today (game is optional)"},
		{"!prac remove <game>", "remove yourself from the prac entry (game is optional)"},
	}
	for _, helper := range helpers {
		response += "**" + helper.Cmd + "** - " + helper.Desc + "\n"
    }
	msg.Respond(response)
}

func (p Prac) fileExists() bool {
	_, err := os.Stat(filePath)
    if err != nil {
        if os.IsNotExist(err) {
			return false
        }
    }
	return true
}

func (p Prac) loadFile() string {
	if !p.fileExists() {
		return ""
	}
	
	// Load raw content of json file
	data, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
	return string(data[:])
}

func (p Prac) saveFile(content jsonContent) {

}

func (p Prac) loadJson() jsonContent {
	var content jsonContent

	fileContent := p.loadFile()
	if len(fileContent) > 0 {
		// Parse file content into json
		dec := json.NewDecoder(strings.NewReader(fileContent))
		for {
			if err := dec.Decode(&content); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
	}

	return content
}

func (p Prac) dateExists(date string, pracJson jsonContent) bool {
	if _, ok := pracJson[date]; ok {
		return true
	}
	return false
}

func (p Prac) findDate(date string, pracJson jsonContent) pracDate {
	for d, games := range pracJson {
		if d == date {
			return games
		}
    }
	return p.newDate()
}

func (p Prac) newDate() pracDate {
	pracData := make(pracDate)
	for _, name := range p.games {
		pracData[name] = pracEntries{}
        
    }
	return pracData
}

func New(games []string) Prac {
	return Prac{games}
}
