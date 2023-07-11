package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token           string
	BotPrefix       string
	FriendBirthdays map[string]string

	config *configStruct
)

type configStruct struct {
	Token           string            `json:"Token"`
	BotPrefix       string            `json:"BotPrefix"`
	FriendBirthdays map[string]string `json:"FriendBirthdays"`
}

func ReadConfig() error {
	fmt.Println("Reading config files")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	FriendBirthdays = config.FriendBirthdays
	return nil
}
