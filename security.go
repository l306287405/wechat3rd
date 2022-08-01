package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type ApplyPrivacyInterfaceReq struct {
	ApiName   string   `json:"api_name"`             //申请的 api 英文名，例如wx.choosePoi，严格区分大小写
	Content   string   `json:"content"`              //申请说原因，不超过300个字符；需要以utf-8编码提交，否则会出现审核失败
	UrlList   []string `json:"url_list,omitempty"`   //(辅助网页)例如，上传官网网页链接用于辅助审核
	PicList   []string `json:"pic_list,omitempty"`   //(辅助图片)填写图片的url ，最多10个
	VideoList []string `json:"video_list,omitempty"` //(辅助视频)填写视频的链接 ，最多支持1个；视频格式只支持mp4格式
}

type ApplyPrivacyInterfaceResp struct {
	core.Error
	AuditId int `json:"audit_id"`
}

//申请地理位置接口
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/privacy-api-management/applyPrivacyInterface.html
func (s *Server) ApplyPrivacyInterface(authorizerAccessToken string, req *ApplyPrivacyInterfaceReq) (resp *ApplyPrivacyInterfaceResp) {
	var (
		u = WECHAT_API_URL + "/wxa/security/apply_privacy_interface?"
	)

	resp = &ApplyPrivacyInterfaceResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetPrivacyInterfaceResp struct {
	core.Error
	InterfaceList []*struct {
		ApiName    string `json:"api_name"`              //api 英文名
		ApiChName  string `json:"api_ch_name"`           //api 中文名
		ApiDesc    string `json:"api_desc"`              //api描述
		ApplyTime  int64  `json:"apply_time,omitempty"`  //申请时间 ，该字段发起申请后才会有
		Status     int    `json:"status,omitempty"`      //接口状态，该字段发起申请后才会有
		AuditId    int    `json:"audit_id,omitempty"`    //申请单号，该字段发起申请后才会有
		FailReason string `json:"fail_reason,omitempty"` //申请被驳回原因或者无权限，该字段申请驳回时才会有
		ApiLink    string `json:"api_link"`              //api文档链接
		GroupName  string `json:"group_name"`            //分组名
	} `json:"interface_list"`
}

//获取地理位置接口列表
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/privacy-api-management/getPrivacyInterface.html
func (s *Server) GetPrivacyInterface(authorizerAccessToken string) (resp *GetPrivacyInterfaceResp) {
	var (
		u = WECHAT_API_URL + "/wxa/security/get_privacy_interface?"
	)
	resp = &GetPrivacyInterfaceResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}
