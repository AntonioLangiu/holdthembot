package bot

import (
	"github.com/AntonioLangiu/holdthembot/common"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

type BotContext struct {
	Config    *common.Configuration
	Bot       *tgbotapi.BotAPI
	UpChannel <-chan tgbotapi.Update
}

func LoadBot(configuration *common.Configuration) {
	ctx := InitBot(configuration)
	RouteMessages(ctx)
}

func InitBot(configuration *common.Configuration) *BotContext {
	ctx := BotContext{}
	ctx.Config = configuration

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramAPI)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	ctx.Bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ctx.UpChannel, err = bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	return &ctx
}
