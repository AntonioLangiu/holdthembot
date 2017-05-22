package bot

import (
	"github.com/AntonioLangiu/holdthembot/common"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
    "os"
)

type BotContext struct {
	Config    *common.Configuration
	Bot       *tgbotapi.BotAPI
	UpChannel <-chan tgbotapi.Update
}

func LoadBot(configuration *common.Configuration) {
	ctx := initBot(configuration)
    initFolder(ctx)
	RouteMessages(ctx)
}

func initFolder(ctx *BotContext) {
    if _, err := os.Stat(ctx.Config.TempFolder); ! os.IsNotExist(err) {
        os.RemoveAll(ctx.Config.TempFolder)
    }
    err := os.Mkdir(ctx.Config.TempFolder, 0777)
    if err != nil {
        log.Fatal(err)
    }
}

func initBot(configuration *common.Configuration) *BotContext {
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
