package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
)

type PluginReq struct {
	Action      string  `json:"action"`
	PluginAppId *string `json:"plugin_appid,omitempty"`
	UserVersion *string `json:"user_version,omitempty"`
}

type PluginResp struct {
	core.Error
	PluginList []*struct {
		AppId      string `json:"appid"`
		Status     int8   `json:"status"`
		Nickname   string `json:"nickname"`
		HeadImgUrl string `json:"headimgurl"`
	} `json:"plugin_list,omitempty"`
}

//小程序插件管理
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Business/pluginManager.html
func (s *Server) Plugin(authorizerAccessToken string, req *PluginReq) (resp *PluginResp) {
	var (
		u = WECHAT_API_URL + "/wxa/plugin?"
	)
	resp = &PluginResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
