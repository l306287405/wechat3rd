package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"log"
	"strconv"
)

type GetcategoryResp struct {
	core.Error
	Data []*Category `json:"data"` //类目信息列表
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//获取当前帐号所设置的类目信息
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/get_category.html
func (s *Server) GetCategory(accessToken string) (resp *GetcategoryResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxaapi/newtmpl/getcategory?"
	)
	resp = &GetcategoryResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

type GetPubTemplateTitlesReq struct {
	Ids   string `json:"ids"`   //类目 id，多个用逗号隔开，可通过接口获取当前帐号所设置的类目信息获取
	Start int    `json:"start"` //用于分页，表示从 start 开始。从 0 开始计数。
	Limit int    `json:"limit"` //用于分页，表示拉取 limit 条记录。最大为 30。
}

type GetPubTemplateTitlesResp struct {
	core.Error
	Count int              `json:"count"`
	Data  []*CategoryTitle `json:"data"`
}

type CategoryTitle struct {
	Tid        int    `json:"tid"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	CategoryId string `json:"categoryId"`
}

//获取模板标题列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/library_list.html
func (s *Server) GetPubTemplateTitles(accessToken string, req *GetPubTemplateTitlesReq) (resp *GetPubTemplateTitlesResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxaapi/newtmpl/getpubtemplatetitles?"
		v = core.AuthTokenUrlValues(accessToken)
	)
	resp = &GetPubTemplateTitlesResp{}
	v.Set("ids", req.Ids)
	v.Set("start", strconv.Itoa(req.Start))
	v.Set("limit", strconv.Itoa(req.Limit))

	err = core.GetRequest(u, v, resp)
	return
}

type GetPubTemplateKeywordsResp struct {
	core.Error
	Data []*CategoryKeyword `json:"data"`
}

type CategoryKeyword struct {
	Kid     int    `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

//获取模板标题下的关键词库
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/library_get.html
func (s *Server) GetPubTemplateKeywords(accessToken string, tid int) (resp *GetPubTemplateKeywordsResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxaapi/newtmpl/getpubtemplatekeywords?"
		v = core.AuthTokenUrlValues(accessToken)
	)
	resp = &GetPubTemplateKeywordsResp{}
	v.Set("tid", strconv.Itoa(tid))

	err = core.GetRequest(u, v, resp)
	return
}

type AddTemplateReq struct {
	Tid       string  `json:"tid"`                 //模板标题 id，可通过获取模板标题列表接口获取，也可登录小程序后台查看获取
	KidList   []int   `json:"kidList"`             //开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
	SceneDesc *string `json:"sceneDesc,omitempty"` //可选,服务场景描述，15个字以内
}

type AddTemplateResp struct {
	core.Error
	PriTmplId string `json:"priTmplId"`
}

//组合模板并添加到个人模板库
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/add_template.html
func (s *Server) AddTemplate(accessToken string, req *AddTemplateReq) (resp *AddTemplateResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxaapi/newtmpl/addtemplate?"
	)
	resp = &AddTemplateResp{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

type GetTemplateResp struct {
	core.Error
	Data []*PriTemplate `json:"data"`
}

type PriTemplate struct {
	PriTmplId string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int    `json:"type"`
}

//获取帐号下的模板列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/list_template.html
func (s *Server) GetTemplate(accessToken string) (resp *GetTemplateResp, err error) {
	var (
		u = WECHAT_API_URL + "/wxaapi/newtmpl/gettemplate?"
	)
	resp = &GetTemplateResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

//删除帐号下的某个模板
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/del_template.html
func (s *Server) DelTemplate(accessToken, priTmplId string) (resp *core.Error, err error) {
	var (
		u   = WECHAT_API_URL + "/wxaapi/newtmpl/deltemplate?"
		req = &struct {
			PriTmplId string `json:"priTmplId"`
		}{PriTmplId: priTmplId}
	)
	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

type sendMsg struct {
	Value string `json:"value"`
}

type SubscribeSendReq struct {
	Touser           string            `json:"touser"`
	TemplateId       string            `json:"template_id"`
	Data             map[string]string `json:"-"`
	Page             *string           `json:"page,omitempty,omitempty"`
	MiniProgramState *string           `json:"miniprogram_state,omitempty"`
	Lang             *string           `json:"lang,omitempty"`

	//请勿填写该参数,并无视它
	DataSource map[string]*sendMsg `json:"data"`
}

//发送订阅消息
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_template/subscribeMessage.send.html
func (s *Server) SubscribeSend(accessToken string, req *SubscribeSendReq) (resp *core.Error, err error) {
	var (
		u = CGIUrl + "/message/subscribe/send?"
	)
	resp = &core.Error{}
	if req.Data!=nil{
		req.DataSource=make(map[string]*sendMsg)

		for k,v:=range req.Data{
			req.DataSource[k]= &sendMsg{Value: v}
			log.Println("k:",k," v:",v," mixed:",*req.DataSource[k])
		}
	}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}
