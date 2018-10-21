package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type Config struct {
    Token string `json:"Token"`
	CommandTrigger string `json:"CommandTrigger"`
	DisabledCmds []string `json:"DisabledCmds"`
    Admins []string `json:"Admins"`
}

func Load() Config {
    var config Config
    configFile, err := os.Open("config/config.json")
    defer configFile.Close()
    if err != nil {
        fmt.Println("error", err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config
}