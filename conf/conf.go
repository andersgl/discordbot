package conf

import (
	"os"
	"encoding/json"
	"fmt"
)

type Conf struct {
    Token string `json:"Token"`
	CommandTrigger string `json:"CommandTrigger"`
	DisabledCmds []string `json:"DisabledCmds"`
}

func Load() Conf {
    var config Conf
    configFile, err := os.Open("conf/config.json")
    defer configFile.Close()
    if err != nil {
        fmt.Println("error", err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config
}