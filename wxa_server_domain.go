package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type ModifyWxaServerDomainReq struct {
	Action                    string  `json:"action"`                                 //get set add delete
	IsModifyPublishedTogether *bool   `json:"is_modify_published_together,omitempty"` //是否同时修改“全网发布版本的值”。（false：只改“测试版”；true：同时改“测试版”和“全网发布版”）省略时，默认为false。
	WxaServerDomain           *string `json:"wxa_server_domain,omitempty"`            //最多可以添加1000个服务器域名，以;隔开。注意：域名不需带有http: // 等协议内容，也不能在域名末尾附加详细的 URI 地址，严格按照类似 www.qq.com 的写法。
}

type ModifyWxaServerDomainResp struct {
	core.Error
	PublishedWxaServerDomain *string `json:"published_wxa_server_domain,omitempty"` //目前生效的 “全网发布版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回。如果没有已发布的第三方平台，该字段也不会返回。
	TestingWxaServerDomain   *string `json:"testing_wxa_server_domain,omitempty"`   //目前生效的 “测试版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回
	InvalidWxaServerDomain   *string `json:"invalid_wxa_server_domain,omitempty"`   //未通过验证的域名。如果不存在未通过验证的域名，该字段不会返回。
}

// 设置第三方平台服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/domain/modify_server_domain.html
func (s *Server) ModifyWxaServerDomain(req *ModifyWxaServerDomainReq) (resp *ModifyWxaServerDomainResp) {
	var (
		u = CGIUrl + "/component/modify_wxa_server_domain?"
	)
	token, err := s.Token()
	resp = &ModifyWxaServerDomainResp{}
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}

type GetDomainConfirmFileResp struct {
	core.Error
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

// 获取第三方业务域名的校验文件
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/domain/get_domain_confirmfile.html
func (s *Server) GetDomainConfirmFile() (resp *GetDomainConfirmFileResp) {
	var (
		u   = CGIUrl + "/component/get_domain_confirmfile?"
		req = &struct{}{}
	)
	token, err := s.Token()
	resp = &GetDomainConfirmFileResp{}
	if err != nil {
		resp.Err(err)
		return
	}
	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}

type ModifyWxaJumpDomainReq struct {
	Action                    string  `json:"action"`                                 //操作类型。可选值请看下文
	IsModifyPublishedTogether *bool   `json:"is_modify_published_together,omitempty"` //是否同时修改“全网发布版本的值”。（false：只改“测试版”；true：同时改“测试版”和“全网发布版”）省略时，默认为false。
	WxaJumpH5Domain           *string `json:"wxa_jump_h5_domain,omitempty"`           //最多可以添加200个小程序业务域名，以;隔开。注意：域名不需带有http:// 等协议内容，也不能在域名末尾附加详细的 URI 地址，严格按照类似 www.qq.com 的写法。
}

type ModifyWxaJumpDomainResp struct {
	core.Error
	PublishedWxaJumpH5Domain *string `json:"published_wxa_jump_h5_domain,omitempty"` //目前生效的 “全网发布版”第三方平台“小程序业务域名”。如果修改失败，该字段不会返回。如果没有已发布的第三方平台，该字段也不会返回。
	TestingWxaJumpH5Domain   *string `json:"testing_wxa_jump_h5_domain,omitempty"`   //目前生效的 “测试版”第三方平台“小程序业务域名”。如果修改失败，该字段不会返回
	InvalidWxaJumpH5Domain   *string `json:"invalid_wxa_jump_h5_domain,omitempty"`   //未通过验证的域名。如果不存在未通过验证的域名，该字段不会返回。
}

// 设置第三方平台业务域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/domain/modify_jump_domain.html
func (s *Server) ModifyWxaJumpDomain(req *ModifyWxaJumpDomainReq) (resp *ModifyWxaJumpDomainResp) {
	var (
		u = CGIUrl + "/component/modify_wxa_jump_domain?"
	)
	token, err := s.Token()
	resp = &ModifyWxaJumpDomainResp{}
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}
