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

// 获取基本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html
func (s *Server) GetAccountBasicInfo(authorizerAccessToken string) (resp *GetAccountBasicInfoResp) {
	var (
		u = CGIUrl + "/account/getaccountbasicinfo?"
	)
	resp = &GetAccountBasicInfoResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

type OpenHaveResp struct {
	core.Error
	HaveOpen bool `json:"have_open"`
}

// 查询公众号/小程序是否绑定open帐号
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/getbindopeninfo.html
func (s *Server) OpenHave(authorizerAccessToken string) (resp *OpenHaveResp) {
	var (
		u = CGIUrl + "/open/have?"
	)
	resp = &OpenHaveResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
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

// 设置服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Server_Address_Configuration.html
func (s *Server) ModifyDomain(authorizerAccessToken string, req *ModifyDomainReq) (resp *ModifyDomainResp) {
	var (
		u = WECHAT_API_URL + "/wxa/modify_domain?"
	)
	resp = &ModifyDomainResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type SetWebviewDomainReq struct {
	Action        *string  `json:"action,omitempty"`
	Webviewdomain []string `json:"webviewdomain,omitempty"`
}

type SetWebviewDomainResp struct {
	core.Error
	Webviewdomain []string `json:"webviewdomain"`
}

// 设置业务域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setwebviewdomain.html
func (s *Server) SetWebviewDomain(authorizerAccessToken string, req *SetWebviewDomainReq) (resp *SetWebviewDomainResp) {
	var (
		u = WECHAT_API_URL + "/wxa/setwebviewdomain?"
	)
	resp = &SetWebviewDomainResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type SetNicknameReq struct {
	Nickname          string  `json:"nick_name"`
	IdCard            *string `json:"id_card,omitempty"`
	License           *string `json:"license,omitempty"`
	NamingOtherStuff1 *string `json:"naming_other_stuff_1,omitempty"`
	NamingOtherStuff2 *string `json:"naming_other_stuff_2,omitempty"`
	NamingOtherStuff3 *string `json:"naming_other_stuff_3,omitempty"`
	NamingOtherStuff4 *string `json:"naming_other_stuff_4,omitempty"`
	NamingOtherStuff5 *string `json:"naming_other_stuff_5,omitempty"`
}

type SetNicknameResp struct {
	core.Error
	Wording string `json:"wording"`
	AuditId int    `json:"auditId"`
}

// 设置名称
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/setnickname.html
func (s *Server) SetNickname(authorizerAccessToken string, req *SetNicknameReq) (resp *SetNicknameResp) {
	var (
		u = WECHAT_API_URL + "/wxa/setwebviewdomain?"
	)
	resp = &SetNicknameResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type CheckWxVerifyNicknameResp struct {
	core.Error
	HitCondition bool   `json:"hit_condition"`
	Wording      string `json:"wording"`
}

// 微信认证名称检测
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/wxverify_checknickname.html
func (s *Server) CheckWxVerifyNickname(authorizerAccessToken string, Nickname string) (resp *CheckWxVerifyNicknameResp) {
	var (
		u   = CGIUrl + "/wxverify/checkwxverifynickname?"
		req = &struct {
			Nickname string `json:"nick_name"`
		}{Nickname: Nickname}
	)
	resp = &CheckWxVerifyNicknameResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type QueryNicknameResp struct {
	core.Error
	Nickname   string `json:"nickname"`
	AuditStat  int8   `json:"audit_stat"`
	FailReason string `json:"fail_reason"`
	CreateTime int64  `json:"create_time"`
	AuditTime  int64  `json:"audit_time"`
}

// 查询改名审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/api_wxa_querynickname.html
func (s *Server) QueryNickname(authorizerAccessToken string, auditId int) (resp *QueryNicknameResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/api_wxa_querynickname?"
		req = &struct {
			AuditId int `json:"audit_id"`
		}{AuditId: auditId}
	)
	resp = &QueryNicknameResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type ModifyHeadImageReq struct {
	HeadImgMediaId string `json:"head_img_media_id"`
	X1             string `json:"x1"`
	Y1             string `json:"y1"`
	X2             string `json:"x2"`
	Y2             string `json:"y2"`
}

// 修改头像
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifyheadimage.html
func (s *Server) ModifyHeadImage(authorizerAccessToken string, req *ModifyHeadImageReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/account/modifyheadimage?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

// 修改简介
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifysignature.html
func (s *Server) ModifySignature(authorizerAccessToken string, signature string) (resp *core.Error) {
	var (
		u   = CGIUrl + "/account/modifysignature?"
		req = &struct {
			Signature string `json:"signature"`
		}{Signature: signature}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetWxaSearchStatusResp struct {
	core.Error
	Status int8 `json:"status"`
}

// 查询搜索设置
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/getwxasearchstatus.html
func (s *Server) GetWxaSearchStatus(authorizerAccessToken string) (resp *GetWxaSearchStatusResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getwxasearchstatus?"
	)
	resp = &GetWxaSearchStatusResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// 修改搜索设置
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/changewxasearchstatus.html
func (s *Server) ChangeWxaSearchStatus(authorizerAccessToken string, status int8) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/changewxasearchstatus?"
		req = &struct {
			Status int8 `json:"status"`
		}{Status: status}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type FetchDataSettingGetResp struct {
	core.Error
	IsPreFetchOpen     bool   `json:"is_pre_fetch_open"`
	PreFetchType       int8   `json:"pre_fetch_type"`
	PreFetchUrl        string `json:"pre_fetch_url,omitempty"`
	PreEnv             string `json:"pre_env"`
	PreFunctionName    string `json:"pre_function_name"`
	IsPeriodFetchOpen  bool   `json:"is_period_fetch_open"`
	PeriodFetchType    int8   `json:"period_fetch_type"`
	PeriodFetchUrl     string `json:"period_fetch_url"`
	PeriodEnv          string `json:"period_env"`
	PeriodFunctionName string `json:"period_function_name"`
}

// 获取数据拉取配置
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/get_fetchdatasetting.html
func (s *Server) FetchDataSettingGet(authorizerAccessToken string) (resp *FetchDataSettingGetResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/fetchdatasetting?"
		req = &struct {
			Action string `json:"action"`
		}{Action: "get"}
	)
	resp = &FetchDataSettingGetResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type FetchDataSettingSetPreFetchReq struct {
	IsPreFetchOpen  bool   `json:"is_pre_fetch_open"`
	PreFetchType    int8   `json:"pre_fetch_type"`
	PreFetchUrl     string `json:"pre_fetch_url,omitempty"`
	PreEnv          string `json:"pre_env"`
	PreFunctionName string `json:"pre_function_name"`
	//IsPeriodFetchOpen  bool   `json:"is_period_fetch_open"`
	//PeriodFetchType    int8   `json:"period_fetch_type"`
	//PeriodFetchUrl     string `json:"period_fetch_url"`
	//PeriodEnv          string `json:"period_env"`
	//PeriodFunctionName string `json:"period_function_name"`
}
type fetchDataSettingSetPreFetchReqWithAction struct {
	Action string `json:"action"`
	FetchDataSettingSetPreFetchReq
}

// 设置预拉取数据
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/set_pre_fetchdatasetting.html
func (s *Server) FetchDataSettingSetPreFetch(authorizerAccessToken string, req *FetchDataSettingSetPreFetchReq) (resp *core.Error) {
	var (
		u  = WECHAT_API_URL + "/wxa/fetchdatasetting?"
		fr = &fetchDataSettingSetPreFetchReqWithAction{Action: "set_pre_fetch", FetchDataSettingSetPreFetchReq: *req}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), fr, resp))
	return
}

type FetchDataSettingSetPeriodFetchReq struct {
	IsPeriodFetchOpen  bool   `json:"is_period_fetch_open"`
	PeriodFetchType    int8   `json:"period_fetch_type"`
	PeriodFetchUrl     string `json:"period_fetch_url,omitempty"`
	PeriodEnv          string `json:"period_env,omitempty"`
	PeriodFunctionName string `json:"period_function_name,omitempty"`
}
type fetchDataSettingSetPeriodFetchReqWithAction struct {
	Action string `json:"action"`
	FetchDataSettingSetPeriodFetchReq
}

// 设置周期性拉取数据
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/set_period_fetchdatasetting.html
func (s *Server) FetchDataSettingSetPeriodFetch(authorizerAccessToken string, req *FetchDataSettingSetPeriodFetchReq) (resp *core.Error) {
	var (
		u  = WECHAT_API_URL + "/wxa/fetchdatasetting?"
		fr = &fetchDataSettingSetPeriodFetchReqWithAction{Action: "set_period_fetch", FetchDataSettingSetPeriodFetchReq: *req}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), fr, resp))
	return
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type GetWxaCodeUnLimitReq struct {
	Scene      string `json:"scene"`                 //最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Page       string `json:"page,omitempty"`        //默认值:主页, 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /, 不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	CheckPath  bool   `json:"check_path,omitempty"`  //默认值:false, 检查 page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但 page 有数量上限（60000个）请勿滥用
	EnvVersion string `json:"env_version,omitempty"` //默认值:release, 要打开的小程序版本。正式版为 release，体验版为 trial，开发版为 develop
	Width      int    `json:"width,omitempty"`       //默认值:430 二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor  bool   `json:"auto_color,omitempty"`  //默认值:false 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColor  *RGB   `json:"line_color,omitempty"`  //auto_color 为 false 时生效，使用 rgb 设置颜色 例如 { "r":"xxx", "g":"xxx", "b":"xxx"} 十进制表示
	IsHyaLine  bool   `json:"is_hya_line,omitempty"` //默认值:false 是否需要透明底色，为 true 时，生成透明底色的小程序
}

type GetWxaCodeUnLimitResp struct {
	core.Error
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

// 获取unlimited小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (s *Server) GetWxaCodeUnLimit(authorizerAccessToken string, req *GetWxaCodeUnLimitReq) (resp *GetWxaCodeUnLimitResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getwxacodeunlimit?"
	)
	resp = &GetWxaCodeUnLimitResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetWxaCodeReq struct {
	Path       string `json:"path"`                  //默认值:主页, 必须是已经发布的小程序存在的页面（否则报错）;最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	EnvVersion string `json:"env_version,omitempty"` //默认值:release, 要打开的小程序版本。正式版为 release，体验版为 trial，开发版为 develop
	Width      int    `json:"width,omitempty"`       //默认值:430 二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor  bool   `json:"auto_color,omitempty"`  //默认值:false 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColor  *RGB   `json:"line_color,omitempty"`  //auto_color 为 false 时生效，使用 rgb 设置颜色 例如 { "r":"xxx", "g":"xxx", "b":"xxx"} 十进制表示
	IsHyaLine  bool   `json:"is_hya_line,omitempty"` //默认值:false 是否需要透明底色，为 true 时，生成透明底色的小程序
}

type GetWxaCodeResp struct {
	core.Error
	Buffer      []byte `json:"buffer"`
	ContentType string `json:"contentType"`
}

// 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (s *Server) GetWxaCode(authorizerAccessToken string, req *GetWxaCodeReq) (resp *GetWxaCodeResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getwxacode?"
	)
	resp = &GetWxaCodeResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
