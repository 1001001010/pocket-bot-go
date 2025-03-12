package main

import (
	"log"

	"github.com/1001001010/pocket-bot-go/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7698482568:AAEOh1oovkYVRXpqbhkBk_IfS6wg2KGKYQA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true


	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}