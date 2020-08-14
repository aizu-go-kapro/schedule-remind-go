package service

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"io/ioutil"
	"net/http"
	"schedule-remind-go/entity"
	"schedule-remind-go/repoository"
)

type calendarAuthServiceImpl struct {
	config *oauth2.Config
	userRepository repoository.UserRepository
}

func GetCalendarAuthService(credentialFileName string, rep repoository.UserRepository) (CalendarAuthService, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}
	service := calendarAuthServiceImpl{
		config:         config,
		userRepository: rep,
	}
	return service, nil
}

func (calendar calendarAuthServiceImpl) GetAuthUrl() string {
	return calendar.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (calendar calendarAuthServiceImpl) VerifyAuthCode(lineId string, authCode string) (*oauth2.Token, error) {
	ctx := context.Background()
	token, err := calendar.config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		LineID:               lineId,
		CalendarAccessToken:  token.AccessToken,
		CalendarTokenType:    token.TokenType,
		CalendarRefreshToken: token.RefreshToken,
		CalendarExpiry:       token.Expiry,
	}
	_, err =calendar.userRepository.Store(ctx, &user)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (calendar calendarAuthServiceImpl) GetAuthToken(lineId string) (*oauth2.Token, error) {
	ctx := context.Background()
	user, err := calendar.userRepository.UserGetByID(ctx, lineId)
	if err != nil {
		return nil, err
	}
	token := oauth2.Token{
		AccessToken:  user.CalendarAccessToken,
		TokenType:    user.CalendarTokenType,
		RefreshToken: user.CalendarRefreshToken,
		Expiry:       user.CalendarExpiry,
	}
	return &token, nil
}

func (calendar calendarAuthServiceImpl) GetClient(lineId string) (*http.Client, error) {
	token, err := calendar.GetAuthToken(lineId)
	if err != nil {
		return nil, err
	}
	return calendar.config.Client(context.Background(), token), nil
}
