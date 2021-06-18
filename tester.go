package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type BindTesterReq struct {
	Wechatid string `json:"wechatid"`
}

type BindTesterResp struct {
	core.Error
	Userstr string `json:"userstr"`
}

//设置业务域名
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html
func (s *Server) BindTester(accessToken string, req *BindTesterReq) (resp *BindTesterResp) {
	var (
		u = WECHAT_API_URL + "/wxa/bind_tester?"
	)
	resp = &BindTesterResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type UnbindTesterReq struct {
	Wechatid *string `json:"wechatid,omitempty"`
	Userstr  *string `json:"userstr,omitempty"`
}

//解除绑定体验者
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html
func (s *Server) UnbindTester(accessToken string, req *UnbindTesterReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/unbind_tester?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type MemberAuthItem struct {
	Userstr string `json:"userstr"`
}

type MemberAuthResp struct {
	core.Error
	Members []*MemberAuthItem `json:"members"`
}

//获取体验者列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/memberauth.html
func (s *Server) MemberAuth(accessToken string) (resp *MemberAuthResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/memberauth?"
		req = &struct {
			Action string `json:"action"`
		}{Action: "get_experiencer"}
	)
	resp = &MemberAuthResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}
