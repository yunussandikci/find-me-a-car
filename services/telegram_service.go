package services

import (
	"find-me-a-car/common"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramService struct {
	bot *tgbotapi.BotAPI
	chatID int64
}

func NewTelegramService(token string, chatID int64) *TelegramService {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return &TelegramService{
		bot: bot,
		chatID: chatID,
	}
}

func (t *TelegramService) SendMessage(message string) {
	msg := tgbotapi.NewMessage(t.chatID, message)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := t.bot.Send(msg)
	if err != nil {
		common.Logger.Errorf(err.Error())
	}
}
