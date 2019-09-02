package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot := buildBot()
	updates, err := bot.GetUpdatesChan(buildUpdateConfig())
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("UserName [%s] Text %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(msg)
		log.Println(err)
	}
}

func buildBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("PORA_ZRAT_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func buildUpdateConfig() tgbotapi.UpdateConfig {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return u
}
