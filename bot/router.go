package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

func RouteMessages(ctx *BotContext) {
	for update := range ctx.UpChannel {
        if update.Message == nil {
			continue
		}
        if update.Message.IsCommand() {
			routeCommand(update, ctx)
        } else if update.Message.VideoNote != nil {
            log.Print("video note!!!!\n\n\n\n")
        } else if update.Message.Voice != nil {
            log.Print("audio!!!!\n\n\n\n")
		} else if update.Message.Text == "" {
			routeInvalid(update, ctx)
		}
	}
}

func routeInvalid(update tgbotapi.Update, ctx *BotContext) {
    HelpMessage(update, ctx)
}

func routeCommand(update tgbotapi.Update, ctx *BotContext) {
	command := strings.ToLower(update.Message.Command())
	var response tgbotapi.MessageConfig
	switch command {
	case "start":
		StartMessage(update, ctx)
	case "help":
		HelpMessage(update, ctx)
	default:
        out := "Sorry, this command is not supported! Send /help to understand how this bot works!"
		response = tgbotapi.NewMessage(update.Message.Chat.ID, out)
	}
	if response.Text != "" {
        SendText(update, ctx, response.Text)
	}
}
