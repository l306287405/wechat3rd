package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"net/url"
)

const WeAPPRegisterUrl = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?"

type FastRegisterWeappReq struct {
	Name               string `json:"name"`                 //企业名
	Code               string `json:"code"`                 //企业代码
	CodeType           int8   `json:"code_type"`            //企业代码类型（1：统一社会信用代码， 2：组织机构代码，3：营业执照注册号）
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
	ComponentPhone     string `json:"component_phone"`      //第三方联系电话
}

//快速注册小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Fast_Registration_Interface_document.html
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
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Fast_Registration_Interface_document.html
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
