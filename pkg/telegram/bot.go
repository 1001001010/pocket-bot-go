package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	poketClient *pocket.Client
	redirectURL string
}

func NewBot(bot *tgbotapi.BotAPI, poketClient *pocket.Client, redirectURL string) *Bot {
	return &Bot{bot: bot, poketClient: poketClient, redirectURL: redirectURL}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}
