package bot

import (
	"botfromlecture/internal/models"
	"log"

	// "math/rand"
	// "strings"
	"time"

	"gopkg.in/telebot.v3"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *models.UserModel
}

var gameItems = [3]string{
	"камень",
	"ножницы",
	"бумага",
}

// var winSticker = &telebot.Sticker{
// 	File: telebot.File{
// 		FileID: "CAACAgIAAxkBAAEGXH5ja6xCMvICwZ3gT4aQZ9vitAEtGAAC3QoAAoeRGEpe4j1vh6BUoCsE",
// 	},
// 	Width:    512,
// 	Height:   512,
// 	Animated: true,
// }

// var loseSticker = &telebot.Sticker{
// 	File: telebot.File{
// 		FileID: "CAACAgIAAxkBAAEGXHxja6vC1RI6t8FD91BtvBiYgZVgtwACHAADwpwNAAFzQzRA6lUbQCsE",
// 	},
// 	Width:    512,
// 	Height:   512,
// 	Animated: true,
// }

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)

		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName + " теперь вы подписаны на рассылку.")
}

// func (bot *UpgradeBot) GameHandler(ctx telebot.Context) error {
// 	return ctx.Send("Сыграем в камень-ножницы-бумага " +
// 		"Введи твой вариант в формате /try камень")
// }

// func (bot *UpgradeBot) TryHandler(ctx telebot.Context) error {
// 	attempts := ctx.Args()

// 	if len(attempts) == 0 {
// 		return ctx.Send("Вы не ввели ваш вариант")
// 	}

// 	if len(attempts) > 1 {
// 		return ctx.Send("Вы ввели больше одного варианта")
// 	}

// 	try := strings.ToLower(attempts[0])
// 	botTry := gameItems[rand.Intn(len(gameItems))]

// 	if botTry == "камень" {
// 		switch try {
// 		case "ножницы":
// 			ctx.Send(loseSticker.Send)
// 			ctx.Send("🪨")
// 			return ctx.Send("Камень! Ты проиграл!")
// 		case "бумага":
// 			ctx.Send(winSticker.Send)
// 			ctx.Send("🪨")
// 			return ctx.Send("Камень! Ты выиграл!")
// 		}
// 	}

// 	if botTry == "ножницы" {
// 		switch try {
// 		case "камень":
// 			ctx.Send(winSticker.Send)
// 			ctx.Send("✂️")
// 			return ctx.Send("Ножницы! Ты выиграл!")
// 		case "бумага":
// 			ctx.Send(loseSticker.Send)
// 			ctx.Send("✂️")
// 			return ctx.Send("Ножницы! Ты проиграл!")
// 		}
// 	}

// 	if botTry == "бумага" {
// 		switch try {
// 		case "ножницы":
// 			ctx.Send(winSticker.Send)
// 			ctx.Send("📃")
// 			return ctx.Send("Бумага! Ты выиграл!")
// 		case "камень":
// 			ctx.Send(loseSticker.Send)
// 			ctx.Send("📃")
// 			return ctx.Send("Бумага! Ты проиграл!")
// 		}
// 	}

// 	if botTry == try {
// 		return ctx.Send("Ничья!")
// 	}

// 	return ctx.Send("Кажется вы ввели неверный вариант!")
// }

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
