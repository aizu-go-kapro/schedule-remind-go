package service

import (
    "github.com/line/line-bot-sdk-go/linebot"
    "net/http"
)

func handleEvent(client *linebot.Client, callbackUrl string, handler func (client *linebot.Client, event *linebot.Event)) {
    http.HandleFunc(callbackUrl, func(writer http.ResponseWriter, request *http.Request) {
        events, err := client.ParseRequest(request)
        if err != nil {
            if err == linebot.ErrInvalidSignature {
                writer.WriteHeader(400)
            } else {
                writer.WriteHeader(500)
            }
        }
        for _, event := range events {
            handler(client, event)
        }
    })
}