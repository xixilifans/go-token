package service

import (
	"context"
	"errors"
	"gologin/model"
	"net/http"
)
//grant 许可
var (
	ErrNotSupportGrantType = errors.New("grant type is not supported")
)

type TokenGranter interface {
	Grant(ctx context.Context,grantType string,client *model.ClientDetails,reader *http.Request) (*model.Oauth2Token,error)
}

type ComposeTokenGranter struct {
	TokenGrantDict map[string]TokenGranter
}


func (tokenGranter *ComposeTokenGranter)Grant(ctx context.Context,grantType string,client *model.ClientDetails,reader *http.Request) ( *model.Oauth2Token,error) {
	//获取具体授权 TokenGranter 生成访问令牌
	dispatchGranter := tokenGranter.TokenGrantDict[grantType]
	if dispatchGranter == nil {
		return nil,ErrNotSupportGrantType
	}
	return dispatchGranter.Grant(ctx,grantType,client,reader)
}

type UsernamePasswordTokenGranter struct {
	supportGrantType string
	userDetailsService UserDetailsService
	tokenService TokenService
}


func NewUsernamePasswordTokenGranter(grantType string,userDetailsService UserDetailsService,tokenService TokenService) TokenGranter {
	return &UsernamePasswordTokenGranter{
		supportGrantType: grantType,
		userDetailsService: userDetailsService,
		tokenService: tokenService,
	}
}

func (tokenGrant *UsernamePasswordTokenGranter) Grant(ctx context.Context,grantType string,client *ClientDetails,reader *http.Request) (*OAuth2Token,error) {
	if grantType != tokenGrant.supportGrantType {
		return nil,ErrNotSupportGrantType
	}
	username := reader.FormValue("username")
	password := reader.FormValue("password")

	if username == "" || password == "" {
		return nil,ErrInvalidUsernameAndPasswordRequest
	}
	//验证用户名密码是否正确
	userDetails,err := tokenGranter.userDetailsService.GetUserDetailsByUsername(ctx,user)
}


type TokenService struct{
	// 根据访问令牌获取对应的用户信息和客户端信息
	GetOAuto2DetailsByAccessToken(tokenValue string) (*OAuth2Details,error)
	// 根据用户信息和客户端信息生成访问令牌
	CreateAccessToken(oauth2Details *OAutn2Details) (*OAuth2Token,error)
	// 根据刷新令牌获取访问令牌
	RefreshAccessToken(refreshTokenValue string) (*OAuth2Token, error)
	// 根据用户信息和客户端信息获取已生成访问令牌
	GetAccessToken(details *OAuth2Details) (*OAuth2Token, error)
	// 根据访问令牌值获取访问令牌结构体
	ReadAccessToken(tokenValue string) (*OAuth2Token, error)
}

