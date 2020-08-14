package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"schedule-remind-go/service"
)

func main() {
	bot, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	service.HandleLineEvent(bot, "/callback", func(client *linebot.Client, event *linebot.Event) {
		// イベントごとの処理
	})
	http.ListenAndServe(":8080", nil)
}
