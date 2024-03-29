package main

import (
	"fmt"

	"github.com/DivyarajChudasama/Discord-bot/bot"
	"github.com/DivyarajChudasama/Discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
