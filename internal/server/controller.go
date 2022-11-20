package server

import (
	"encoding/json"
	"io"
	"log"
	"nazirov-tagaev-fp-bot/cmd/bot"
	"net/http"
)

type msgReq struct {
	MsgId uint64 `json:"msg_id"`
}

type Bot struct {
	BotInstance *bot.MailingBot
}

func (b *Bot) sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var Msg msgReq

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Cannot read request body: %v", err)
	}

	err = json.Unmarshal(reqBody, &Msg)
	if err != nil {
		log.Fatalf("Cannot unmarshal json body from request: %v", err)
	}

	b.BotInstance.SendMailing(Msg.MsgId)

	_, err = io.WriteString(w, "Message received!")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func InitServer(bot *bot.MailingBot) {
	existBot := Bot{bot}

	mux := http.NewServeMux()
	mux.HandleFunc("/message", existBot.sendMessageHandler)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
