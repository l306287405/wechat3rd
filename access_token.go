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
	var(
		ticket string
		resp *AccessTokenResponse
	)

	if d.ExpiresIn <= time.Now().Unix()-30 {
		ticket, err = d.GetTicket()
		if err != nil {
			return
		}
		resp, err = newAccessToken(&AccessTokenRequest{
			ComponentAppid:        d.AppID,
			ComponentAppsecret:    d.AppSecret,
			ComponentVerifyTicket: ticket,
		})
		if err != nil {
			return
		}
		d.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
		d.ComponentAccessToken = resp.ComponentAccessToken
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
func newAccessToken(r *AccessTokenRequest) (*AccessTokenResponse, error) {
	resp := &AccessTokenResponse{}
	err := core.PostJson(componentAccessTokenUrl, r, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}