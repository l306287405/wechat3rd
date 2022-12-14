package wechat3rd

import (
	"errors"
	"fmt"
	"github.com/l306287405/wechat3rd/cache"
	"sync"
	"time"

	"github.com/l306287405/wechat3rd/core"
)

const (
	componentAccessTokenUrl = WECHAT_API_URL + "/cgi-bin/component/api_component_token"
)

type AccessTokenServer interface {
	Token() (token string, err error)
	SetToken(token string, expiresIn int64) (err error)
	RefreshToken() (token string, err error)
}

type DefaultAccessTokenServer struct {
	mux   sync.Mutex
	Cache cache.Cache
	TicketServer
	AppID     string
	AppSecret string
	Cfg       Config
	refresh   bool
}

// 获取令牌   token不使用不获取
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func (d *DefaultAccessTokenServer) Token() (token string, err error) {

	if d.refresh {
		return d.token()
	}
	if cacheToken := d.Cache.Get(d.getCacheKey()); cacheToken != nil {
		return cacheToken.(string), nil
	}
	return d.token()
}

func (d *DefaultAccessTokenServer) token() (token string, err error) {
	var (
		resp *AccessTokenResp
	)

	ticket, err := d.GetTicket()
	if err != nil {
		return
	}

	d.mux.Lock()
	defer d.mux.Unlock()

	resp = newAccessToken(&AccessTokenReq{
		ComponentAppid:        d.AppID,
		ComponentAppsecret:    d.AppSecret,
		ComponentVerifyTicket: ticket,
	})
	if !resp.Success() {
		err = errors.New(fmt.Sprintf("get component_access_token errcode: %d,errmsg: %s", resp.ErrCode, resp.ErrMsg))
		return
	}

	_ = d.SetToken(resp.ComponentAccessToken, resp.ExpiresIn)

	return resp.ComponentAccessToken, nil
}

// 从别处恢复token
func (d *DefaultAccessTokenServer) RefreshToken() (token string, err error) {
	d.refresh = true
	token, err = d.Token()
	d.refresh = false
	return
}

// 获取令牌（component_access_token),带有过期时间的时间戳。
func (d *DefaultAccessTokenServer) SetToken(token string, expiresIn int64) (err error) {
	return d.Cache.Set(d.getCacheKey(), token, time.Duration(expiresIn-120)*time.Second)
}

func (d *DefaultAccessTokenServer) getCacheKey() string {
	return "wechat_open_platform.access_token." + d.Cfg.AppID
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
