package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type BindTesterReq struct {
	Wechatid string `json:"wechatid"`
}

type BindTesterResp struct {
	core.Error
	Userstr string `json:"userstr"`
}

//绑定微信用户为体验者
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html
func (s *Server) BindTester(authorizerAccessToken string, req *BindTesterReq) (resp *BindTesterResp) {
	var (
		u = WECHAT_API_URL + "/wxa/bind_tester?"
	)
	resp = &BindTesterResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type UnbindTesterReq struct {
	Wechatid *string `json:"wechatid,omitempty"`
	Userstr  *string `json:"userstr,omitempty"`
}

//解除绑定体验者
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html
func (s *Server) UnbindTester(authorizerAccessToken string, req *UnbindTesterReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/unbind_tester?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
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
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/memberauth.html
func (s *Server) MemberAuth(authorizerAccessToken string) (resp *MemberAuthResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/memberauth?"
		req = &struct {
			Action string `json:"action"`
		}{Action: "get_experiencer"}
	)
	resp = &MemberAuthResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
