package model

//定义客户端信息

type ClientDetails struct {
	//client 标识
	ClientId string
	//client密钥
	ClientSecret string
	//访问Token 有效时间
	AccessTokenValiditySecond int
	// 刷新令牌有效时间，秒
    RefreshTokenValiditySeconds int
    // 重定向地址，授权码类型中使用
    RegisteredRedirectUri string
    // 可以使用的授权类型
    AuthorizedGrantTypes []string


}

