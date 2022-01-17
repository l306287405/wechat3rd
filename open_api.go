package wechat3rd

import "github.com/l306287405/wechat3rd/core"

//清空api的调用quota
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/openApi/clear_quota.html
func (s *Server) ClearQuota(authorizerAccessToken string, appId string) (resp *core.Error) {
	var (
		u   = CGIUrl + "/clear_quota?"
		req = &struct {
			AppId string `json:"appid"`
		}{AppId: appId}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type QuotaGetResp struct {
	core.Error
	Quota struct {
		DailyLimit int `json:"daily_limit"`
		Used       int `json:"used"`
		Remain     int `json:"remain"`
	} `json:"quota"`
}

//查询openApi调用quota
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/openApi/get_api_quota.html
func (s *Server) QuotaGet(authorizerAccessToken string, cgiPath string) (resp *QuotaGetResp) {
	var (
		u   = CGIUrl + "/openapi/quota/get?"
		req = &struct {
			CgiPath string `json:"cgi_path"`
		}{CgiPath: cgiPath}
	)
	resp = &QuotaGetResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type RidGetResp struct {
	core.Error
	Request struct {
		InvokeTime   int64  `json:"invoke_time"`
		CostInMs     int    `json:"cost_in_ms"`
		RequestUrl   string `json:"request_url"`
		RequestBody  string `json:"request_body"`
		ResponseBody string `json:"response_body"`
	} `json:"request"`
}

//查询rid信息
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/openApi/get_rid_info.html
func (s *Server) RidGet(authorizerAccessToken string, rid string) (resp *RidGetResp) {
	var (
		u   = CGIUrl + "/openapi/rid/get?"
		req = &struct {
			Rid string `json:"rid"`
		}{Rid: rid}
	)
	resp = &RidGetResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
