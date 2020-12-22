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
	)
	resp = &Jscode2sessionResp{}
	req.Set("appid",appId)
	req.Set("js_code",jsCode)
	req.Set("grant_type","authorization_code")
	req.Set("component_appid",s.cfg.AppID)
	req.Set("component_access_token",s.getToken())

	err=core.GetRequest(u,req,resp)
	return
}

type UserInfoResp struct {
	Subscribe int8 `json:"subscribe"`
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int8 `json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	SubscribeTime string `json:"subscribeTime"`
	Unionid string `json:"unionid"`
	Remark string `json:"remark"`
	Groupid int `json:"groupid"`
	TagidList []int `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene int `json:"qr_scene"`
	QrSceneStr int `json:"qr_scene_str"`
}

func (s *Server) UserInfo(openId string)(resp *UserInfoResp,err error) {
	var(
		u = "https://api.weixin.qq.com/cgi-bin/user/info?"
		p = make(url.Values)
	)
	resp = &UserInfoResp{}
	p.Set("access_token",s.getToken())
	p.Set("openid",openId)
	err=core.GetRequest(u,p,resp)
	return
}