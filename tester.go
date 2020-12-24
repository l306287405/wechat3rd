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
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/setwebviewdomain.html
func (s *Server) BindTester(authToken string,req *BindTesterReq) (resp *BindTesterResp,err error){
	var(
		u = wechatApiUrl+"/wxa/bind_tester?"
	)
	resp = &BindTesterResp{}

	err=core.PostJson(s.AuthToken2url(u,authToken),req,resp)
	return
}

type UnbindTesterReq struct {
	Wechatid *string `json:"wechatid,omitempty"`
	Userstr *string `json:"userstr,omitempty"`
}

//解除绑定体验者
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html
func (s *Server) UnbindTester(authToken string,req *UnbindTesterReq) (resp *core.Error,err error){
	var(
		u = wechatApiUrl+"/wxa/unbind_tester?"
	)
	resp = &core.Error{}

	err=core.PostJson(s.AuthToken2url(u,authToken),req,resp)
	return
}

type MemberAuthResp struct {
	core.Error
	Members []struct{
		Userstr string `json:"userstr"`
	} `json:"members"`
}

//解除绑定体验者
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html
func (s *Server) MemberAuth(authToken string) (resp *MemberAuthResp,err error){
	var(
		u = wechatApiUrl+"/wxa/memberauth?"
		req = &struct {
			Action string
		}{Action: "get_experiencer"}
	)
	resp = &MemberAuthResp{}
	err=core.PostJson(s.AuthToken2url(u,authToken),req,resp)
	return
}