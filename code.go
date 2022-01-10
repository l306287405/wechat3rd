package wechat3rd

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/l306287405/wechat3rd/core"
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

type CommitReq struct {
	TemplateId  int    `json:"template_id"`  //代码库中的代码模板 ID
	ExtJson     string `json:"ext_json"`     //第三方自定义的配置
	UserVersion string `json:"user_version"` //代码版本号，开发者可自定义（长度不要超过 64 个字符）
	UserDesc    string `json:"user_desc"`    //代码描述，开发者可自定义
}

//上传小程序代码
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html
func (s *Server) Commit(accessToken string, req *CommitReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/commit?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type GetPageResp struct {
	core.Error
	PageList []string `json:"page_list"` //page_list 页面配置列表
}

//获取已上传的代码的页面列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (s *Server) GetPage(accessToken string) (resp *GetPageResp) {
	var (
		u = WECHAT_API_URL + "/wxa/get_page?"
	)
	resp = &GetPageResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//获取体验版二维码 参数path 为官方参数 ,参数saveDir为二维码图片存储路径 参数fileName 为二维码图片存储名称请勿包含名称
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (s *Server) GetQrcode(accessToken string, path, saveDir, fileName *string) (filePath string, err error) {
	var (
		u        = WECHAT_API_URL + "/wxa/get_qrcode?"
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

	if saveDir == nil {
		saveDir = new(string)
		*saveDir = "/var/tmp/" + httpResp.Header.Get("Content-Type")
	}
	*saveDir = strings.TrimRight(*saveDir, "/")
	_, err = os.Stat(*saveDir) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		if err = os.MkdirAll(*saveDir, 0755); err != nil {
			return
		}
	}

	filePath = *saveDir + "/" + *fileName + ".jpg"

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
	AuditId int `json:"auditid"`
}

//提交审核
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (s *Server) SubmitAudit(accessToken string, req *SubmitAuditReq) (resp *SubmitAuditResp) {
	var (
		u = WECHAT_API_URL + "/wxa/submit_audit?"
	)
	resp = &SubmitAuditResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type GetAuditStatusResp struct {
	core.Error
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	Screenshot string `json:"screenshot"`
}

//查询指定发布审核单的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html
func (s *Server) GetAuditStatus(accessToken string, auditId int) (resp *GetAuditStatusResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/get_auditstatus?"
		req = &struct {
			AuditId int `json:"auditid"`
		}{AuditId: auditId}
	)
	resp = &GetAuditStatusResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type GetLatestAuditStatusResp struct {
	core.Error
	AuditId    int    `json:"auditid"`
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"screen_shot"`
}

//查询最新一次提交的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_latest_auditstatus.html
func (s *Server) GetLatestAuditStatus(accessToken string) (resp *GetLatestAuditStatusResp) {
	var (
		u = WECHAT_API_URL + "/wxa/get_latest_auditstatus?"
	)
	resp = &GetLatestAuditStatusResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//小程序审核撤回
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/undocodeaudit.html
func (s *Server) UndoCodeAudit(accessToken string) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/undocodeaudit?"
	)
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//发布已通过审核的小程序
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html
func (s *Server) Release(accessToken string) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/release?"
		req = &struct{}{}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

//版本回退
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertcoderelease.html
func (s *Server) RevertCodeRelease(accessToken string, appVersion int) (resp *core.Error) {
	var (
		u      = WECHAT_API_URL + "/wxa/revertcoderelease?"
		params = core.AuthTokenUrlValues(accessToken)
	)
	//版本为零则回退到上一个版本
	if appVersion != 0 {
		params.Set("app_version", strconv.Itoa(appVersion))
	}
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, params, resp))
	return
}

type GetRevertCodeReleaseResp struct {
	core.Error
	TemplateList []*RevertTemplate `json:"version_list"` //版本信息列表
}

type RevertTemplate struct {
	CommitTime  int64  `json:"commit_time"`  //更新时间，时间戳
	UserVersion string `json:"user_version"` //版本号，开发者自定义字段
	UserDesc    string `json:"user_desc"`    //版本描述   开发者自定义字段
	AppVersion  int    `json:"app_version"`  //小程序版本
}

//获取可回退的小程序版本
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_history_version.html
func (s *Server) GetRevertCodeRelease(accessToken string) (resp *GetRevertCodeReleaseResp) {
	var (
		u      = WECHAT_API_URL + "/wxa/revertcoderelease?"
		params = core.AuthTokenUrlValues(accessToken)
	)
	params.Set("action", "get_history_version")
	resp = &GetRevertCodeReleaseResp{}
	resp.Err(core.GetRequest(u, params, resp))
	return
}

type GetPaidUnionIdReq struct {
	OpenId        string  `json:"openid"`
	TransactionId *string `json:"transaction_id,omitempty"`
	MchId         *string `json:"mch_id,omitempty"`
	OutTradeNo    *string `json:"out_trade_no,omitempty"`
}

type GetPaidUnionIdResp struct {
	core.Error
	UnionId string `json:"unionid,omitempty"`
}

//支付后获取用户 Unionid 接口
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/User_Management.html
func (s *Server) GetPaidUnionId(accessToken string, req *GetPaidUnionIdReq) (resp *GetPaidUnionIdResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getpaidunionid?"
		p = make(url.Values)
	)
	resp = &GetPaidUnionIdResp{}

	p.Set("openid", req.OpenId)

	if req.TransactionId != nil {
		p.Set("transaction_id", *req.TransactionId)
	} else {
		if req.MchId == nil && req.OutTradeNo == nil {
			resp.Err(errors.New("参数缺失"))
			return
		}
		p.Set("mch_id", *req.MchId)
		p.Set("out_trade_no", *req.OutTradeNo)
	}
	u += p.Encode()

	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//分阶段发布
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/grayrelease.html
func (s *Server) GrayRelease(accessToken string, grayPercentage int8) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/grayrelease?"
		req = &struct {
			GrayPercentage int8 `json:"gray_percentage"`
		}{GrayPercentage: grayPercentage}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type GetGrayReleasePlanResp struct {
	core.Error
	GrayReleasePlan *GrayReleasePlan `json:"gray_release_plan"` //模板信息列表
}

type GrayReleasePlan struct {
	Status          int8  `json:"status"`           //0:初始状态 1:执行中 2:暂停中 3:执行完毕 4:被删除
	CreateTimestamp int64 `json:"create_timestamp"` //分阶段发布计划的创建事件
	GrayPercentage  int8  `json:"gray_percentage"`  //当前的灰度比例
}

//查询当前分阶段发布详情
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html
func (s *Server) GetGrayReleasePlan(accessToken string) (resp *GetGrayReleasePlanResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getgrayreleaseplan?"
	)
	resp = &GetGrayReleasePlanResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//取消分阶段发布
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertgrayrelease.html
func (s *Server) RevertGrayRelease(accessToken string) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/revertgrayrelease?"
	)
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//修改小程序服务状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/change_visitstatus.html
func (s *Server) ChangeVisitStatus(accessToken string, action string) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/change_visitstatus?"
		req = &struct {
			Action string `json:"action"`
		}{Action: action}
	)

	if action != "open" && action != "close" {
		resp.Err(errors.New("action must be 'open' or 'close'"))
		return
	}

	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type GetWeappSupportVersionResp struct {
	core.Error
	NowVersion string    `json:"now_version"` //当前版本
	UvInfo     *struct { //版本的用户占比列表
		Items []*struct {
			Percentage float64 `json:"percentage"` //百分比
			Version    string  `json:"version"`    //基础库版本号
		} `json:"items"`
	} `json:"uv_info"`
}

//查询当前设置的最低基础库版本及各版本用户占比
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getweappsupportversion.html
func (s *Server) GetWeappSupportVersion(accessToken string) (resp *GetWeappSupportVersionResp) {
	var (
		u   = CGIUrl + "/wxopen/getweappsupportversion?"
		req = &struct{}{}
	)

	resp = &GetWeappSupportVersionResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

//设置最低基础库版本
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/setweappsupportversion.html
func (s *Server) SetWeappSupportVersion(accessToken string, version string) (resp *core.Error) {
	var (
		u   = CGIUrl + "/wxopen/setweappsupportversion?"
		req = &struct {
			Version string `json:"version"`
		}{Version: version}
	)

	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type QueryQuotaResp struct {
	core.Error
	Rest         int `json:"rest"`          //quota剩余值
	Limit        int `json:"limit"`         //当月分配quota
	SpeedupRest  int `json:"speedup_rest"`  //剩余加急次数
	SpeedupLimit int `json:"speedup_limit"` //当月分配加急次数
}

//查询服务商的当月提审限额（quota）和加急次数
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/query_quota.html
func (s *Server) QueryQuota(accessToken string) (resp *QueryQuotaResp) {
	var (
		u = WECHAT_API_URL + "/wxa/queryquota?"
	)
	resp = &QueryQuotaResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(accessToken), resp))
	return
}

//加急审核申请
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html
func (s *Server) SpeedupAudit(accessToken string, auditId int) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/speedupaudit?"
		req = &struct {
			AuditId int `json:"auditid"`
		}{AuditId: auditId}
	)

	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}
