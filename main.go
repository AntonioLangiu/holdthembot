package main

import (
	"github.com/AntonioLangiu/holdthembot/bot"
	"github.com/AntonioLangiu/holdthembot/common"
)

func main() {
	configuration := common.LoadConfiguration()
	go common.WebServer()
	bot.LoadBot(configuration)
}
