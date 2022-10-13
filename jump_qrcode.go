package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type GetJumpQRCodeReq struct {
	Appid      string   `json:"appid"`
	GetType    int      `json:"get_type"`
	PrefixList []string `json:"prefix_list"`
	PageNum    int      `json:"page_num"`
	PageSize   int      `json:"page_size"`
}
type GetJumpQRCodeResp struct {
	core.Error
	QrcodejumpOpen     int     `json:"qrcodejump_open"`
	ListSize           int     `json:"list_size"`
	QrcodejumpPubQuota int     `json:"qrcodejump_pub_quota"`
	TotalCount         int     `json:"total_count"`
	RuleList           []*Rule `json:"rule_list"`
}

type Rule struct {
	Prefix      string   `json:"prefix"`
	OpenVersion int      `json:"open_version"`
	State       int      `json:"state"`
	Path        string   `json:"path"`
	DebugUrl    []string `json:"debug_url"`
}

//获取已设置的二维码规则
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/getJumpQRCode.html
func (s *Server) GetJumpQRCode(authorizerAccessToken string, req *GetJumpQRCodeReq) (resp *GetJumpQRCodeResp) {
	var (
		u = CGIUrl + "/wxopen/qrcodejumpget?"
	)
	resp = &GetJumpQRCodeResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type AddJumpQRCodeReq struct {
	IsEdit        int      `json:"is_edit"`
	Prefix        string   `json:"prefix"`
	Path          string   `json:"path"`
	OpenVersion   int      `json:"open_version"`
	DebugUrl      []string `json:"debug_url"`
	PermitSubRule int      `json:"permit_sub_rule"`
	Appid         string   `json:"appid"`
}

//增加或修改二维码规则
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/addJumpQRCode.html
func (s *Server) AddJumpQRCode(authorizerAccessToken string, req *AddJumpQRCodeReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/wxopen/qrcodejumpadd?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type PublishJumpQRCodeReq struct {
	Prefix string `json:"prefix"`
}

//发布已设置的二维码规则
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/publishJumpQRCode.html
func (s *Server) PublishJumpQRCode(authorizerAccessToken string, req *PublishJumpQRCodeReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/wxopen/qrcodejumppublish?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type DeleteJumpQRCodeReq struct {
	Prefix string `json:"prefix"`
	Appid  string `json:"appid"`
}

//删除已设置的二维码规则
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/deleteJumpQRCode.html
func (s *Server) DeleteJumpQRCode(authorizerAccessToken string, req *DeleteJumpQRCodeReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/wxopen/qrcodejumpdelete?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type DownloadQRCodeTextResp struct {
	core.Error
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

//获取二维码规则校验文件名称及内容
//https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/downloadQRCodeText.html
func (s *Server) DownloadQRCodeText(authorizerAccessToken string) (resp *DownloadQRCodeTextResp) {
	var (
		u   = CGIUrl + "/wxopen/qrcodejumpdownload?"
		req = &struct{}{}
	)
	resp = &DownloadQRCodeTextResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
