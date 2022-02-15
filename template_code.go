package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type GetTemplateDraftListResp struct {
	core.Error
	DraftList []*Draft `json:"draft_list"` //草稿信息列表
}

type Draft struct {
	CreateTime  int64  `json:"create_time"`  //开发者上传草稿时间戳
	UserVersion string `json:"user_version"` //版本号，开发者自定义字段
	UserDesc    string `json:"user_desc"`    //版本描述   开发者自定义字段
	DraftId     int    `json:"draft_id"`     //草稿 id
}

//获取代码草稿列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (s *Server) GetTemplateDraftList() (resp *GetTemplateDraftListResp) {
	var (
		u     = WECHAT_API_URL + "/wxa/gettemplatedraftlist?"
		token string
		err   error
	)
	resp = &GetTemplateDraftListResp{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(token), resp))
	return
}

//将草稿添加到代码模板库
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html
func (s *Server) AddToTemplate(draftId int) (resp *core.Error) {
	var (
		u     = WECHAT_API_URL + "/wxa/addtotemplate?"
		token string
		req   = &struct {
			DraftId int `json:"draft_id"`
		}{DraftId: draftId}
		err error
	)
	resp = &core.Error{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}

type GetTemplateListResp struct {
	core.Error
	TemplateList []*Template `json:"template_list"` //模板信息列表
}

type Template struct {
	CreateTime  int64  `json:"create_time"`  //被添加为模板的时间
	UserVersion string `json:"user_version"` //版本号，开发者自定义字段
	UserDesc    string `json:"user_desc"`    //版本描述   开发者自定义字段
	TemplateId  int    `json:"template_id"`  //模板 id
}

//获取代码模板列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html
func (s *Server) GetTemplateList() (resp *GetTemplateListResp) {
	var (
		u     = WECHAT_API_URL + "/wxa/gettemplatelist?"
		token string
		err   error
	)
	resp = &GetTemplateListResp{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(token), resp))
	return
}

//删除指定代码模板
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html
func (s *Server) DeleteTemplate(templateId int) (resp *core.Error) {
	var (
		u     = WECHAT_API_URL + "/wxa/deletetemplate?"
		token string
		req   = &struct {
			TemplateId int `json:"template_id"`
		}{TemplateId: templateId}
		err error
	)
	resp = &core.Error{}
	token, err = s.Token()
	if err != nil {
		resp.Err(err)
		return
	}

	resp.Err(core.PostJson(s.AuthToken2url(u, token), req, resp))
	return
}
