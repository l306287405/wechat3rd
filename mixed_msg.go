package wechat3rd

type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`

	//代码审核结果推送  weapp_audit_success
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
	Event        string `xml:"Event" json:"Event"`
	SuccTime     int64  `xml:"SuccTime" json:"SuccTime"`

	//代码审核结果推送 weapp_audit_fail
	CodeAuditReason string `xml:"Reason" json:"Reason"`
	FailTime        int64  `xml:"FailTime" json:"FailTime"`
	ScreenShot      string `xml:"ScreenShot" json:"ScreenShot"`

	//代码审核结果推送 weapp_audit_delay
	DelayTime int64 `xml:"DelayTime" json:"DelayTime"`

	//echo
	EchoStr string `xml:"-" json:"-"`

	MsgId        int64   `xml:"MsgId"        json:"MsgId"`        // request
	Content      string  `xml:"Content"      json:"Content"`      // request
	MediaId      string  `xml:"MediaId"      json:"MediaId"`      // request
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`       // request
	Format       string  `xml:"Format"       json:"Format"`       // request
	Recognition  string  `xml:"Recognition"  json:"Recognition"`  // request
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"` // request
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`   // request
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`   // request
	Scale        int     `xml:"Scale"        json:"Scale"`        // request
	Label        string  `xml:"Label"        json:"Label"`        // request
	Title        string  `xml:"Title"        json:"Title"`        // request
	Description  string  `xml:"Description"  json:"Description"`  // request
	URL          string  `xml:"Url"          json:"Url"`          // request
	EventKey     string  `xml:"EventKey"     json:"EventKey"`     // request, menu
	Ticket       string  `xml:"Ticket"       json:"Ticket"`       // request
	Latitude     float64 `xml:"Latitude"     json:"Latitude"`     // request
	Longitude    float64 `xml:"Longitude"    json:"Longitude"`    // request
	Precision    float64 `xml:"Precision"    json:"Precision"`    // request
	BizMsgMenuId int64   `xml:"bizmsgmenuid" json:"bizmsgmenuid"` // request
	// menu
	MenuId       int64 `xml:"MenuId" json:"MenuId"`
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item" json:"PicList"`
	} `xml:"SendPicsInfo" json:"SendPicsInfo"`
	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		PoiName   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo" json:"SendLocationInfo"`

	MsgID  int64  `xml:"MsgID"  json:"MsgID"`  // template, mass
	Status string `xml:"Status" json:"Status"` // template, mass
	// shakearound
	ChosenBeacon struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"ChosenBeacon" json:"ChosenBeacon"`
	AroundBeacons []struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"AroundBeacons>AroundBeacon" json:"AroundBeacons"`

	UnionId string `xml:"UnionId"              json:"UnionId"` // unionId

	// component_verify_ticket 验证票据推送
	AppId string `xml:"AppId" json:"AppId"`
	//CreateTime int32 `xml:"CreateTime" json:"CreateTime"`
	InfoType              string `xml:"InfoType" json:"InfoType"`
	MiniAppId             string `xml:"appid" json:"appid"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"`

	// 授权变更通知推送 authorized unauthorized updateauthorized
	AuthorizerAppid              string `xml:"AuthorizerAppid" json:"AuthorizerAppid"`
	AuthorizationCode            string `xml:"AuthorizationCode" json:"AuthorizationCode"`
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime" json:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `xml:"PreAuthCode" json:"PreAuthCode"`

	//notify_third_fasteregister notify_third_fastverifybetaapp 注册审核事件推送 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
	//MiniAppId             string `xml:"appid" json:"appid"`
	EventStatus int    `xml:"status" json:"status"`
	AuthCode    string `xml:"auth_code" json:"auth_code"`
	Msg         string `xml:"msg" json:"msg"`
	Info        struct {
		Name               string `xml:"name" json:"name"`
		Code               string `xml:"code" json:"code"`
		CodeType           int8   `xml:"code_type" json:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat" json:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name" json:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone" json:"component_phone"`
		WxUser             string `xml:"wxuser" json:"wxuser"`
		IdName             string `xml:"idname" json:"idname"`

		//notify_third_fastregisterbetaapp 创建试用小程序成功/失败的通知数据 推送 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html
		UniqueId string `xml:"unique_id" json:"unique_id"`
	} `xml:"info" json:"info"`

	// wxa_nickname_audit 名称审核结果事件推送 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/wxa_nickname_audit.html
	//Ret    int8   `json:"ret" xml:"ret"`
	Nickname string `xml:"nickname" json:"nickname"`
	//Reason string `xml:"reason" json:"reason"`

	// wxa_category_audit 类目推送 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/wxa_category_audit.html
	First  string `xml:"first" json:"first"`
	Second string `xml:"second" json:"second"`
	Ret    int8   `xml:"ret" json:"ret"`
	Reason string `xml:"reason" json:"reason"`

	// 小程序申诉记录推送 wxa_appeal_record https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/records/record_event.html
	AppealRecordId int64 `xml:"appeal_record_id" json:"appeal_record_id"`
	//MiniAppId      string `xml:"appid" json:"appid"`
	AppealTime        int64  `xml:"appeal_time" json:"appeal_time"`
	AppealCount       int    `xml:"appeal_count" json:"appeal_count"`
	AppealFrom        int8   `xml:"appeal_from" json:"appeal_from"`
	AppealStatus      int8   `xml:"appeal_status" json:"appeal_status"`
	AuditTime         int64  `xml:"audit_time" json:"audit_time"`
	AuditReason       string `xml:"audit_reason" json:"audit_reason"`
	PunishDescription string `xml:"punish_description" json:"punish_description"`
	Material          []struct {
		IllegalMaterial struct {
			Content    string `xml:"content" json:"content"`
			ContentUrl string `xml:"content_url" json:"content_url"`
		} `xml:"illegal_material" json:"illegal_material"`
		AppealMaterial struct {
			Reason          string `xml:"reason" json:"reason"`
			ProofMaterialId string `xml:"proof_material_id" json:"proof_material_id"`
		} `xml:"appeal_material" json:"appeal_material"`
	} `xml:"material" json:"material"`

	// wxa_live_apply_event 小程序直播 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Business/live_player/applyliveinfo.html
	ApplyLiveInfoNoitfy struct {
		AppId    string `xml:"appid" json:"appid"`
		OpenTime int64  `xml:"open_time" json:"open_time"`
	} `xml:"ApplyLiveInfoNoitfy" json:"ApplyLiveInfoNoitfy"`

	// wxa_media_check 异步内容安全识别 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
	TraceId string `xml:"trace_id" json:"trace_id"`
	Version int8   `xml:"version" json:"version"`
	Detail  []struct {
		Strategy string `xml:"strategy" json:"strategy"`
		ErrCode  int    `xml:"errcode" json:"errcode"`
		Suggest  string `xml:"suggest" json:"suggest"`
		Label    int8   `xml:"label" json:"label"`
		Prob     int8   `xml:"prob" json:"prob"`
	} `xml:"detail" json:"detail"`
	ErrCode int    `xml:"errcode" json:"errcode"`
	ErrMsg  string `xml:"errmsg" json:"errmsg"`
	Result  struct {
		Suggest string `xml:"suggest" json:"suggest"`
		Label   int8   `xml:"label" json:"label"`
	} `xml:"result" json:"result"`
}
