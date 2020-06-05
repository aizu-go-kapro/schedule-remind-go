package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

type Config struct {
	LineChannelSecret string
	LineChannelAccessToken string
}

var conf Config

func init() {
	c, _ := ini.Load("config.ini")
	conf = Config{
		LineChannelSecret: c.Section("line").Key("channel_secret").String(),
		LineChannelAccessToken: c.Section("line").Key("channel_access_token").String(),
	}
}

func main() {
	bot, err := linebot.New(conf.LineChannelSecret, conf.LineChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {
		events, err := bot.ParseRequest(request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				writer.WriteHeader(400)
			} else {
				writer.WriteHeader(500)
			}
		}
		for _, event := range events{
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
