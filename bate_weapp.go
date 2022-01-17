package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type FastRegisterBetaWeappReq struct {
	Name   string `json:"name"`
	OpenId string `json:"openid"`
}

type FastRegisterBetaWeappResp struct {
	core.Error
	UniqueId     string `json:"unique_id"`
	AuthorizeUrl string `json:"authorize_url"`
}

// 创建试用小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html
func (s *Server) FastRegisterBetaWeapp(req *FastRegisterBetaWeappReq) (resp *FastRegisterBetaWeappResp) {
	var (
		u = WECHAT_API_URL + "/wxa/component/fastregisterbetaweapp?"
	)
	token, err := s.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	resp = &FastRegisterBetaWeappResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}

type VerifyBetaWeappReq struct {
	EnterpriseName     string `json:"enterprise_name"`      //企业名（需与工商部门登记信息一致)；如果是“无主体名称个体工商户”则填“个体户+法人姓名”，例如“个体户张三”
	Code               string `json:"code"`                 //企业代码
	CodeType           int8   `json:"code_type"`            //企业代码类型 1：统一社会信用代码（18 位） 2：组织机构代码（9 位 xxxxxxxx-x） 3：营业执照注册号(15 位)
	LegalPersonaWechat string `json:"legal_persona_wechat"` //法人微信号
	LegalPersonaName   string `json:"legal_persona_name"`   //法人姓名（绑定银行卡）
	LegalPersonaIdCard string `json:"legal_persona_idcard"` //法人身份证号
	ComponentPhone     string `json:"component_phone"`      //第三方联系电话
}

// 试用小程序快速认证
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastverify.html
func (s *Server) VerifyBetaWeapp(accessToken string, req *VerifyBetaWeappReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/verifybetaweapp?"
		r = struct {
			VerifyInfo *VerifyBetaWeappReq `json:"verify_info"`
		}{VerifyInfo: req}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), r, resp))
	return
}

// 修改试用小程序名称
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastmodify.html
func (s *Server) SetBetaWeappNickname(accessToken string, name string) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/setbetaweappnickname?"
		r = struct {
			Name string `json:"name"`
		}{Name: name}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), r, resp))
	return
}
