package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type OpenCreateOrGetResp struct {
	core.Error
	OpenAppId string `json:"open_appid"`
}

// 创建开放平台帐号并绑定公众号/小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/create.html
func (s *Server) OpenCreate(accessToken string, appId *string) (resp *OpenCreateOrGetResp) {
	var (
		u   = CGIUrl + "/open/create?"
		req = &struct {
			AppId *string `json:"appid"`
		}{}
	)
	if appId != nil {
		req.AppId = appId
	}
	resp = &OpenCreateOrGetResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type OpenBindOrUnbindReq struct {
	AppId     *string `json:"appid,omitempty"` //非必填，如果不填则取生成authorizer_access_token的授权公众号或小程序的 appid。如果填，则需要填与生成authorizer_access_token的授权公众号或小程序的 appid一致的appid，否则会出现40013报错
	OpenAppid string  `json:"open_appid"`      //开放平台帐号 appid，由创建开发平台帐号接口返回
}

// 将公众号/小程序绑定到开放平台帐号下
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/create.html
func (s *Server) OpenBind(accessToken string, req *OpenBindOrUnbindReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/open/bind?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

// 将公众号/小程序从开放平台帐号下解绑
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/unbind.html
func (s *Server) OpenUnbind(accessToken string, req *OpenBindOrUnbindReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/open/unbind?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

// 获取公众号/小程序所绑定的开放平台帐号
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/account/get.html
func (s *Server) OpenGet(accessToken string, appId *string) (resp *OpenCreateOrGetResp) {
	var (
		u   = CGIUrl + "/open/get?"
		req = &struct {
			AppId *string `json:"appid,omitempty"`
		}{}
	)
	if appId != nil {
		req.AppId = appId
	}
	resp = &OpenCreateOrGetResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}
