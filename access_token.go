package wechat3rd

import (
	"errors"
	"fmt"
	"time"

	"github.com/l306287405/wechat3rd/core"
)

const (
	componentAccessTokenUrl = WECHAT_API_URL + "/cgi-bin/component/api_component_token"
)

type AccessTokenServer interface {
	Token() (token string, err error)
	GetToken() (token string, expiresIn int64, err error)
	RestoreToken(token string, expiresIn int64) (err error)
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
	var (
		ticket string
		resp   *AccessTokenResp
	)

	if d.ExpiresIn <= time.Now().Unix()-30 {
		ticket, err = d.GetTicket()
		if err != nil {
			return
		}
		resp = newAccessToken(&AccessTokenReq{
			ComponentAppid:        d.AppID,
			ComponentAppsecret:    d.AppSecret,
			ComponentVerifyTicket: ticket,
		})
		if !resp.Success() {
			err = errors.New(fmt.Sprintf("get component_access_token errcode: %d,errmsg: %s", resp.ErrCode, resp.ErrMsg))
			return
		}
		d.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
		d.ComponentAccessToken = resp.ComponentAccessToken
	}
	return d.ComponentAccessToken, nil
}

// 从别处恢复token
func (d *DefaultAccessTokenServer) RestoreToken(token string, expiresIn int64) (err error) {
	d.ExpiresIn = expiresIn
	d.ComponentAccessToken = token
	return
}

// 获取令牌（component_access_token),带有过期时间的时间戳。
func (d *DefaultAccessTokenServer) GetToken() (token string, expiresIn int64, err error) {
	if token, err = d.Token(); err == nil {
		expiresIn = d.ExpiresIn
	}

	return
}

type AccessTokenResp struct {
	core.Error
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

type AccessTokenReq struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// 获取第三方应用token
func newAccessToken(r *AccessTokenReq) *AccessTokenResp {
	resp := &AccessTokenResp{}
	resp.Err(core.PostJson(componentAccessTokenUrl, r, resp))
	return resp
}
