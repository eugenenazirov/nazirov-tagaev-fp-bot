package bot

import (
	"log"
	"nazirov-tagaev-fp-bot/internal/models"
	"time"

	"gopkg.in/telebot.v3"
)

type MailingBot struct {
	Bot      *telebot.Bot
	Users    *models.UserModel
	Messages *models.MessageModel
}

func (bot *MailingBot) SendMailing(msgId uint64) {
	users, err := bot.Users.FindAll()
	if err != nil {
		log.Printf("Ошибка получения пользователей %v", err)
	}

	msgText, err := bot.Messages.FindOne(msgId)
	if err != nil {
		log.Printf("Ошибка получения сообщения %v", err)
	}

	for _, user := range users {
		_, err := bot.Bot.Send(user, msgText.MsgText)
		if err != nil {
			log.Printf("Ошибка получения пользователей %v", err)
		}
	}
}

func (bot *MailingBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Пользователь не получен %v, попробуем его создать", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)

		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName + ", теперь вы подписаны на рассылку.")
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}
