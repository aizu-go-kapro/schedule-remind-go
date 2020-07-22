package entity

import "time"

type User struct {
	LineID string `db:"line_id"`
	CalendarAccessToken string `db:"calendar_access_token"`
	CalendarTokenType string `db:"calendar_token_type"`
	CalendarRefreshToken string `db:"calendar_refresh_token"`
	CalendarExpiry time.Time `db:"calendar_expiry"`
}
