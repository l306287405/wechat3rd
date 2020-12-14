package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"time"
)

const (
	componentAccessTokenUrl = wechatApiUrl + "/cgi-bin/component/api_component_token"
)

type AccessTokenServer interface {
	Token() (token string, err error)
}

type DefaultAccessTokenServer struct {
	TicketServer
	AppID     string
	AppSecret string

	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"` // 当前时间 + 过期时间
}

// token不使用不获取
func (d *DefaultAccessTokenServer) Token() (token string, err error) {
	timeUnix := time.Now().Unix()
	if d.ExpiresIn <= time.Now().Unix()-30 {
		ticket, err := d.GetTicket()
		if err != nil {
			return "", nil
		}
		aresp, err := newAccessToken(d.AppID, d.AppSecret, ticket)
		if err != nil {
			return "", nil
		}
		d.ExpiresIn = timeUnix + aresp.ExpiresIn
		d.ComponentAccessToken = aresp.ComponentAccessToken
	}
	return d.ComponentAccessToken, nil
}

type AccessTokenResponse struct {
	core.Error
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

type AccessTokenRequest struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// 获取第三方应用token
func newAccessToken(appid, AppSecret, ticket string) (*AccessTokenResponse, error) {
	req := AccessTokenRequest{
		ComponentAppid:        appid,
		ComponentAppsecret:    AppSecret,
		ComponentVerifyTicket: ticket,
	}
	resp := &AccessTokenResponse{}
	err := core.PostJson(componentAccessTokenUrl, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}