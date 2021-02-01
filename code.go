package wechat3rd

import (
	"errors"
	"github.com/l306287405/wechat3rd/core"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

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
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (s *Server) GetTemplateDraftList() (resp *GetTemplateDraftListResp, err error) {
	var (
		u     = wechatApiUrl + "/wxa/gettemplatedraftlist?"
		token string
	)
	token, err = s.Token()
	if err != nil {
		return
	}
	resp = &GetTemplateDraftListResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(token), resp)
	return
}

//将草稿添加到代码模板库
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html
func (s *Server) AddToTemplate(draftId int) (resp *core.Error, err error) {
	var (
		u     = wechatApiUrl + "/wxa/addtotemplate?"
		token string
		req   = &struct {
			DraftId int `json:"draft_id"`
		}{DraftId: draftId}
	)
	token, err = s.Token()
	if err != nil {
		return
	}

	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
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
	TemplateId  string `json:"template_id"`  //模板 id
}

//获取代码模板列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (s *Server) GetTemplateList() (resp *GetTemplateListResp, err error) {
	var (
		u     = wechatApiUrl + "/wxa/gettemplatelist?"
		token string
	)
	token, err = s.Token()
	if err != nil {
		return
	}
	resp = &GetTemplateListResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(token), resp)
	return
}

//删除指定代码模板
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (s *Server) DeleteTemplate(templateId string) (resp *core.Error, err error) {
	var (
		u     = wechatApiUrl + "/wxa/deletetemplate?"
		token string
		req   = &struct {
			TemplateId string `json:"template_id"`
		}{TemplateId: templateId}
	)
	token, err = s.Token()
	if err != nil {
		return
	}

	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	return
}

type CommitReq struct {
	TemplateId  string `json:"template_id"`  //代码库中的代码模板 ID
	ExtJson     string `json:"ext_json"`     //第三方自定义的配置
	UserVersion string `json:"user_version"` //代码版本号，开发者可自定义（长度不要超过 64 个字符）
	UserDesc    string `json:"user_desc"`    //代码描述，开发者可自定义
}

//上传小程序代码
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
func (s *Server) Commit(accessToken string, req *CommitReq) (resp *core.Error, err error) {
	var (
		u = wechatApiUrl + "/wxa/commit?"
	)
	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

type GetPageResp struct {
	core.Error
	PageList []string `json:"page_list"` //page_list 页面配置列表
}

//获取已上传的代码的页面列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
func (s *Server) GetPage(accessToken string) (resp *GetPageResp, err error) {
	var (
		u = wechatApiUrl + "/wxa/get_page?"
	)
	resp = &GetPageResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

//获取体验版二维码
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_qrcode.html
func (s *Server) GetQrcode(accessToken string, path *string, saveRoot *string) (filePath string, err error) {
	var (
		u        = wechatApiUrl + "/wxa/get_qrcode?"
		p        = core.AuthTokenUrlValues(accessToken)
		httpResp *http.Response
		fp       *os.File
	)

	if path != nil {
		p.Set("path", url.QueryEscape(*path))
	}
	httpResp, err = http.Get(u + p.Encode())
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = errors.New("http.Status:" + httpResp.Status)
		return
	}

	if saveRoot == nil {
		saveRoot = new(string)
		*saveRoot = "/var/tmp/" + httpResp.Header.Get("Content-Type")
	}
	filePath = *saveRoot + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"

	fp, err = os.Create(filePath)
	if err != nil {
		return
	}
	defer fp.Close()

	_, err = fp.ReadFrom(httpResp.Body)
	return filePath, err
}

type SubmitAuditReq struct {
	ItemList      []*SubmitAuditItem      `json:"item_list,omitempty"`
	PreviewInfo   *SubmitAuditPreviewInfo `json:"preview_info,omitempty"`
	VersionDesc   *string                 `json:"version_desc,omitempty"`
	FeedbackInfo  *string                 `json:"feedback_info,omitempty"`
	FeedbackStuff *string                 `json:"feedback_stuff,omitempty"`
	UgcDeclare    *SubmitAuditUgcDeclare  `json:"ugc_declare,omitempty"`
}

type SubmitAuditUgcDeclare struct {
	Scene          []int8  `json:"scene,omitempty"`
	OtherSceneDesc *string `json:"other_scene_desc,omitempty"`
	Method         []int8  `json:"method,omitempty"`
	HasAuditTeam   *int8   `json:"has_audit_team,omitempty"`
	AuditDesc      *string `json:"audit_desc,omitempty"`
}

type SubmitAuditPreviewInfo struct {
	VideoIdList []string `json:"video_id_list,omitempty"`
	PicIdList   []string `json:"pic_id_list,omitempty"`
}

type SubmitAuditItem struct {
	Address     *string `json:"address,omitempty"`
	Tag         *string `json:"tag,omitempty"`
	FirstClass  *string `json:"first_class,omitempty"`
	SecondClass *string `json:"second_class,omitempty"`
	ThirdClass  *string `json:"third_class,omitempty"`
	FirstId     *int64  `json:"first_id,omitempty"`
	SecondId    *int64  `json:"second_id,omitempty"`
	ThirdId     *int64  `json:"third_id,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type SubmitAuditResp struct {
	core.Error
	AuditId int64 `json:"auditid"`
}

//提交审核
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
func (s *Server) SubmitAudit(accessToken string, req *SubmitAuditReq) (resp *SubmitAuditResp, err error) {
	var (
		u = wechatApiUrl + "/wxa/submit_audit?"
	)
	resp = &SubmitAuditResp{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

type GetAuditStatusResp struct {
	core.Error
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	Screenshot string `json:"screenshot"`
}

//查询指定发布审核单的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html
func (s *Server) GetAuditStatus(accessToken string, auditId string) (resp *GetAuditStatusResp, err error) {
	var (
		u   = wechatApiUrl + "/wxa/get_auditstatus?"
		req = &struct {
			AuditId string `json:"auditid"`
		}{AuditId: auditId}
	)
	resp = &GetAuditStatusResp{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

type GetLatestAuditStatusResp struct {
	core.Error
	AuditId    string `json:"auditid"`
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"screen_shot"`
}

//查询最新一次提交的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_latest_auditstatus.html
func (s *Server) GetLatestAuditStatus(accessToken string) (resp *GetLatestAuditStatusResp, err error) {
	var (
		u = wechatApiUrl + "/wxa/get_latest_auditstatus?"
	)
	resp = &GetLatestAuditStatusResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

//小程序审核撤回
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/undocodeaudit.html
func (s *Server) UndoCodeAudit(accessToken string) (resp *core.Error, err error) {
	var (
		u = wechatApiUrl + "/wxa/undocodeaudit?"
	)
	resp = &core.Error{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

//发布已通过审核的小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/release.html
func (s *Server) Release(accessToken string) (resp *core.Error, err error) {
	var (
		u   = wechatApiUrl + "/wxa/release?"
		req = &struct{}{}
	)
	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, accessToken), req, resp)
	return
}

//版本回退
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertcoderelease.html
func (s *Server) RevertCodeRelease(accessToken string) (resp *core.Error, err error) {
	var (
		u = wechatApiUrl + "/wxa/revertcoderelease?"
	)
	resp = &core.Error{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}

type GetPaidUnionIdReq struct {
	OpenId        string  `json:"openid"`
	TransactionId *string `json:"transaction_id,omitempty"`
	MchId         *string `json:"mch_id,omitempty"`
	OutTradeNo    *string `json:"out_trade_no,omitempty"`
}

type GetPaidUnionIdResp struct {
	UnionId string `json:"unionid,omitempty"`
	ErrCode int    `json:"errcode"`
}

//支付后获取用户 Unionid 接口
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/User_Management.html
func (s *Server) GetPaidUnionId(accessToken string, req *GetPaidUnionIdReq) (resp *GetPaidUnionIdResp, err error) {
	var (
		u = wechatApiUrl + "/wxa/getpaidunionid?"
		p = make(url.Values)
	)
	resp = &GetPaidUnionIdResp{}

	p.Set("openid", req.OpenId)

	if req.TransactionId != nil {
		p.Set("transaction_id", *req.TransactionId)
	} else {
		if req.MchId == nil && req.OutTradeNo == nil {
			err = errors.New("参数缺失")
			return
		}
		p.Set("mch_id", *req.MchId)
		p.Set("out_trade_no", *req.OutTradeNo)
	}
	u += p.Encode()

	err = core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp)
	return
}
