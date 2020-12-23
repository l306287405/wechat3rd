package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"net/url"
)

type Jscode2sessionResp struct {
	Openid string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Errcode int `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}

func (s *Server) Jscode2session(appId ,jsCode string) (resp *Jscode2sessionResp,err error){
	var(
		req = make(url.Values)
		u = "https://api.weixin.qq.com/sns/component/jscode2session?"
		accessToken string
	)
	accessToken,err=s.Token()
	if err!=nil{
		return
	}
	resp = &Jscode2sessionResp{}
	req.Set("appid",appId)
	req.Set("js_code",jsCode)
	req.Set("grant_type","authorization_code")
	req.Set("component_appid",s.cfg.AppID)
	req.Set("component_access_token",accessToken)

	err=core.GetRequest(u,req,resp)
	return
}
