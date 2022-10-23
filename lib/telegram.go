package lib

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var lockTg = &sync.Mutex{}

var tgInstance *tgbotapi.BotAPI

func GetTelegramInstance(apikey string) *tgbotapi.BotAPI {
	if tgInstance == nil {
		lockTg.Lock()
		defer lockTg.Unlock()
		if tgInstance == nil {
			bot, err := tgbotapi.NewBotAPI(apikey)
			if err != nil {
				panic(err)
			}
			tgInstance = bot
		}
	}

	return tgInstance
}
