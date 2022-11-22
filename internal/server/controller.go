package server

import (
	"encoding/json"
	"io"
	"log"
	"nazirov-tagaev-fp-bot/cmd/bot"
	"net/http"
)

// msgReq imitates a request body from admin panel
type msgReq struct {
	MsgId uint64 `json:"msg_id"`
}

// Bot struct links existing Bot instance with sendMessageHandler method
type Bot struct {
	BotInstance *bot.MailingBot
}

// sendMessageHandler parses request body from admin panel
// takes the message id from it and calls
// the bot.SendMailing() method with message id
// as the parameter
// return void
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

// InitServer initializes the web-server instance using ServeMux
// takes the bot instance link as an argument
// return void
func InitServer(bot *bot.MailingBot) {
	existBot := Bot{bot}

	mux := http.NewServeMux()
	mux.HandleFunc("/message", existBot.sendMessageHandler)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
