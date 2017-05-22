package bot

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendText(update tgbotapi.Update, ctx *BotContext, text string) {
	response := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	response.ParseMode = "HTML"
	ctx.Bot.Send(response)
}

func StartMessage(update tgbotapi.Update, ctx *BotContext) {
	out := "Hi! I'm <b>HoldThemBot</b> 🤖  and I will help you to "
    out += "<b>download voice and video messages</b> as files.\n"
    out += "It's easy, just send me audio and video messages and "
    out += "I will send them back to you!\nHave fun! 😊"
    SendText(update, ctx, out)
}

func HelpMessage(update tgbotapi.Update, ctx *BotContext) {
    out := "I will help you to download <b>voice</b> and "
    out += "video messages! Just send them to me and I will "
    out += "send you them back as a file!\n"
    out += "You can also forward voice and video from other "
    out += "chats and get them as a file! Try me, send me a "
    out += "video message now!!🤖"
	SendText(update, ctx, out)
}
