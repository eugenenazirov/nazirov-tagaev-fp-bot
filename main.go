package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"nazirov-tagaev-fp-bot/cmd/bot"
	"nazirov-tagaev-fp-bot/internal/models"
	"nazirov-tagaev-fp-bot/internal/server"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	mailingBot := bot.MailingBot{
		Bot:      bot.InitBot(cfg.BotToken),
		Users:    &models.UserModel{Db: db},
		Messages: &models.MessageModel{Db: db},
	}

	mailingBot.Bot.Handle("/start", mailingBot.StartHandler)

	go server.InitServer(&mailingBot)

	mailingBot.Bot.Start()
}
