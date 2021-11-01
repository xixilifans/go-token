package model

import "time"

type Oauth2Token struct {
	//refresh token
	RefreshToken *Oauth2Token
	//type of token
	TokenType string
	//token
	TokenValue string
	//out of date time
	ExpiresTime *time.Time
}
