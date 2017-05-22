package bot

import (
	"gopkg.in/telegram-bot-api.v4"
)

func SendText(update tgbotapi.Update, ctx *BotContext, text string) {
	response := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	response.ParseMode = "HTML"
	ctx.Bot.Send(response)
}

func StartMessage(update tgbotapi.Update, ctx *BotContext) {
	out := "Ciao! Sono HKNBot, il 🤖  dell'Associazione IEEE-Eta Kappa Nu" +
		" del Politecnico di Torino. Organizziamo Eventi e Gruppi di" +
		" studio, e tramite me potrai avere tutte le informazioni" +
		" di cui hai bisogno 👍  Sono un Bot testuale, per cui scrivimi e" +
		" e cercherò di risponderti!"
    SendText(update, ctx, out)
}

func HelpMessage(update tgbotapi.Update, ctx *BotContext) {
	out := "Sono un bot testuale, basato sull'"
	out += "NLP, meglio conosciuta come "
	out += "elaborazione del linguaggio naturale.\n"
	out += "Per comunicare con me puoi scrivermi frasi e io "
	out += "cercherò di risponderti al meglio! 😊\n"
	SendText(update, ctx, out)
}
