package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"net/url"
)

const OrgWeAPPRegisterUrl = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?"
const PersonalWeAPPRegisterUrl = "https://api.weixin.qq.com/wxa/component/fastregisterpersonalweapp?"

type FastRegisterWeappReq struct {
	Name               string `json:"name"`                 //企业名
	Code               string `json:"code"`                 //企业代码
	CodeType           int8   `json:"code_type"`            //企业代码类型（1：统一社会信用代码， 2：组织机构代码，3：营业执照注册号）
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
	ComponentPhone     string `json:"component_phone"`      //第三方联系电话
}

// FastRegisterWeapp 快速注册小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
//Deprecated: 转用 FastRegisterOrgWeapp 方法
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
	resp.Err(core.PostJson(OrgWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type FastRegisterOrgWeappReq struct {
	Name               string `json:"name"`                 //企业名
	Code               string `json:"code"`                 //企业代码
	CodeType           int8   `json:"code_type"`            //企业代码类型（1：统一社会信用代码， 2：组织机构代码，3：营业执照注册号）
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
	ComponentPhone     string `json:"component_phone"`      //第三方联系电话
}

// FastRegisterOrgWeapp 快速注册小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
func (s *Server) FastRegisterOrgWeapp(req *FastRegisterWeappReq) (resp *core.Error) {
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
	resp.Err(core.PostJson(OrgWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type SearchWeappReq struct {
	Name               string `json:"name"`                 //企业名
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
}

// SearchWeapp 查询创建任务状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
//Deprecated: 转用 QueryOrgWeapp 方法
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
	resp.Err(core.PostJson(OrgWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type QueryOrgWeappReq struct {
	Name               string `json:"name"`                 //企业名
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
}

// QueryOrgWeapp 查询创建任务状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
func (s *Server) QueryOrgWeapp(req *SearchWeappReq) (resp *core.Error) {
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
	resp.Err(core.PostJson(OrgWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type FastRegisterPersonalWeappReq struct {
	IdName         string  `json:"idname"`                    //个人用户名称
	WxUser         string  `json:"wxuser"`                    //个人用户微信号
	ComponentPhone *string `json:"component_phone,omitempty"` //第三方联系电话
}

type FastRegisterPersonalWeappResp struct {
	core.Error
	taskid        string // 任务id，后面query查询需要用到
	authorize_url string // 给用户扫码认证的验证url
	status        int    // 任务的状态
}

// FastRegisterPersonalWeapp 快速注册个人小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html
func (s *Server) FastRegisterPersonalWeapp(req *FastRegisterPersonalWeappReq) (resp *FastRegisterPersonalWeappResp) {
	var (
		p     = make(url.Values)
		token string
		err   error
	)
	resp = &FastRegisterPersonalWeappResp{}

	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	p.Set("action", "create")
	p.Set("component_access_token", token)
	resp.Err(core.PostJson(PersonalWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}

type QueryPersonalWeappResp struct {
	core.Error
	TaskId       string `json:"taskid"`
	AuthorizeUrl string `json:"authorize_url"`
	Status       int    `json:"status"`
}

//查询个人小程序创建任务状态接口详情
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html
func (s *Server) QueryPersonalWeapp(taskId string) (resp *QueryPersonalWeappResp) {
	var (
		p     = make(url.Values)
		token string
		err   error
		req   = &struct {
			TaskId string `json:"taskid"`
		}{TaskId: taskId}
	)
	resp = &QueryPersonalWeappResp{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	p.Set("action", "query")
	p.Set("component_access_token", token)
	resp.Err(core.PostJson(PersonalWeAPPRegisterUrl+p.Encode(), req, resp))
	return
}
