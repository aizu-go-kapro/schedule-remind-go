package service

import (
	"golang.org/x/oauth2"
	"net/http"
)

type CalendarAuthService interface {
	GetAuthUrl() string
	SaveAuthCode(lineId string, authCode string) (*oauth2.Token, error)
	GetAuthToken(lineId string) (*oauth2.Token, error)
	GetClient(lineId string) (*http.Client, error)
}
