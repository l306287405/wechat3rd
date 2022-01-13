package wechat3rd

import (
	"net/url"

	"github.com/l306287405/wechat3rd/core"
)

const (
	WeAPPRegisterUrl = CGIUrl + "/component/fastregisterweapp?"
	FastregisterUrl  = CGIUrl + "/account/fastregister"
	MPAuthUrl        = WECHAT_MP_URL + "/cgi-bin/fastregisterauth"
)

type FastRegisterWeappReq struct {
	Name               string `json:"name"`                 //企业名
	Code               string `json:"code"`                 //企业代码
	CodeType           int8   `json:"code_type"`            //企业代码类型（1：统一社会信用代码， 2：组织机构代码，3：营业执照注册号）
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
	ComponentPhone     string `json:"component_phone"`      //第三方联系电话
}

//快速注册小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
func (s *Server) FastRegisterWeapp(req *FastRegisterWeappReq) (resp *core.Error) {
	var (
		p     = make(url.Values)
		token string
		err   error
	)
	resp = &core.Error{}

	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	p.Set("action", "create")
	p.Set("component_access_token", token)
	resp.Err(core.PostJson(WeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type SearchWeappReq struct {
	Name               string `json:"name"`                 //企业名
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
}

//查询创建任务状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
func (s *Server) SearchWeapp(req *SearchWeappReq) (resp *core.Error) {
	var (
		p     = make(url.Values)
		token string
		err   error
	)
	resp = &core.Error{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	p.Set("action", "search")
	p.Set("component_access_token", token)
	resp.Err(core.PostJson(WeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

//获取微信公众平台授权页面链接
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fast_registration_of_mini_program.html
func (srv *Server) GetMPAuthUrl(mpAppID, redirectUri string, copyWxVerify bool) (u string) {
	p := url.Values{}
	p.Set("component_appid", srv.cfg.AppID)
	p.Set("appid", mpAppID)
	cwv := "1"
	if !copyWxVerify {
		cwv = "0"
	}
	p.Set("copy_wx_verify", cwv)
	p.Set("redirect_uri", redirectUri)
	u = MPAuthUrl + "?" + p.Encode()
	return
}

//复用公众号主体快速注册小程序
type FastRegisterReq struct {
	Ticket string `json:"ticket"` //公众号扫码授权的凭证(公众平台扫码页面回跳到第三方平台时携带)
}
type FastRegisterResp struct {
	core.Error
	Appid             string `json:"appid"`              //新创建小程序的 appid
	AuthorizationCode string `json:"authorization_code"` //新创建小程序的授权码
	IsWxVerifySucc    bool   `json:"is_wx_verify_succ"`  //复用公众号微信认证小程序是否成功
	IsLinkSucc        bool   `json:"is_link_succ"`       //小程序是否和公众号关联成功
}

//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fast_registration_of_mini_program.html
func (s *Server) FastRegister(accessToken string, req *FastRegisterReq) (resp *FastRegisterResp) {
	resp = &FastRegisterResp{}
	tUrl := s.AuthToken2url(FastregisterUrl, accessToken)
	resp.Err(core.PostJson(tUrl, req, resp))
	return
}
