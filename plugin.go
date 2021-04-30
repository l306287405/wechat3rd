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
	PluginList  []*struct{
		AppId  string `json:"appid"`
		Status  int8 `json:"status"`
		Nickname string `json:"nickname"`
		HeadImgUrl	string `json:"headimgurl"`
	} `json:"plugin_list,omitempty"`
}

//小程序插件管理
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Plug-ins_Management.html
func (s *Server) Plugin(accessToken string, req *PluginReq) (resp *PluginResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxa/plugin?"
	)
	resp = &PluginResp{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}
