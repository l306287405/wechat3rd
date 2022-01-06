package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
)

type GetAccountBasicInfoResp struct {
	core.Error
	Appid          string `json:"appid"`
	AccountType    int8   `json:"account_type"`
	PrincipalType  int    `json:"principal_type"`
	PrincipalName  string `json:"principal_name"`
	RealnameStatus int8   `json:"realname_status"`
	WxVerifyInfo   struct {
		QualificationVerify   bool   `json:"qualification_verify"`
		NamingVerify          bool   `json:"naming_verify"`
		AnnualReview          *bool  `json:"annual_review,omitempty"`
		AnnualReviewBeginTime *int64 `json:"annual_review_begin_time,omitempty"`
		AnnualReviewEndTime   *int64 `json:"annual_review_end_time,omitempty"`
	} `json:"wx_verify_info"`
	SignatureInfo struct {
		Signature       string `json:"signature"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"signature_info"`
	HeadImageInfo struct {
		HeadImageUrl    string `json:"head_image_url"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"head_image_info"`
	NicknameInfo struct {
		Nickname        string `json:"nickname"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"nickname_info"`
	RegisteredCountry int `json:"registered_country"`
}

//获取基本信息
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html
func (s *Server) GetAccountBasicInfo(authToken string) (resp *GetAccountBasicInfoResp) {
	var (
		u = CGIUrl + "/account/getaccountbasicinfo?"
	)
	resp = &GetAccountBasicInfoResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authToken), resp))
	return
}

type ModifyDomainReq struct {
	Action          string   `json:"action"`
	Requestdomain   []string `json:"requestdomain"`
	Wsrequestdomain []string `json:"wsrequestdomain"`
	Uploaddomain    []string `json:"uploaddomain"`
	Downloaddomain  []string `json:"downloaddomain"`
	Udpdomain       []string `json:"udpdomain"`
	Tcpdomain       []string `json:"tcpdomain"`
}

type ModifyDomainResp struct {
	core.Error
	Requestdomain          []string `json:"requestdomain"`
	Wsrequestdomain        []string `json:"wsrequestdomain"`
	Uploaddomain           []string `json:"uploaddomain"`
	Downloaddomain         []string `json:"downloaddomain"`
	Udpdomain              []string `json:"udpdomain"`
	Tcpdomain              []string `json:"tcpdomain"`
	InvalidRequestdomain   []string `json:"invalid_requestdomain"`
	InvalidWsrequestdomain []string `json:"invalid_wsrequestdomain"`
	InvalidUploaddomain    []string `json:"invalid_uploaddomain"`
	InvalidDownloaddomain  []string `json:"invalid_downloaddomain"`
	InvalidUdpdomain       []string `json:"invalid_udpdomain"`
	InvalidTcpdomain       []string `json:"invalid_tcpdomain"`
	NoIcpDomain            []string `json:"no_icp_domain"`
}

//设置服务器域名
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Server_Address_Configuration.html
func (s *Server) ModifyDomain(authToken string, req *ModifyDomainReq) (resp *ModifyDomainResp) {
	var (
		u = WECHAT_API_URL + "/wxa/modify_domain?"
	)
	resp = &ModifyDomainResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authToken), req, resp))
	return
}

type SetWebviewDomainReq struct {
	Action        *string  `json:"action,omitempty"`
	Webviewdomain []string `json:"webviewdomain,omitempty"`
}

//设置业务域名
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setwebviewdomain.html
func (s *Server) SetWebviewDomain(authToken string, req *SetWebviewDomainReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/setwebviewdomain?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authToken), req, resp))
	return
}
