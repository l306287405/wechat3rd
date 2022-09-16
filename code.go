package wechat3rd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/l306287405/wechat3rd/core"
)

type CommitReq struct {
	TemplateId  int    `json:"template_id"`  //代码库中的代码模板 ID
	ExtJson     string `json:"ext_json"`     //第三方自定义的配置
	UserVersion string `json:"user_version"` //代码版本号，开发者可自定义（长度不要超过 64 个字符）
	UserDesc    string `json:"user_desc"`    //代码描述，开发者可自定义
}

// Commit 上传小程序代码
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html
func (s *Server) Commit(authorizerAccessToken string, req *CommitReq) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/commit?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetPageResp struct {
	core.Error
	PageList []string `json:"page_list"` //page_list 页面配置列表
}

// GetPage 获取已上传的代码的页面列表
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (s *Server) GetPage(authorizerAccessToken string) (resp *GetPageResp) {
	var (
		u = WECHAT_API_URL + "/wxa/get_page?"
	)
	resp = &GetPageResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// GetQrcode 获取体验版二维码 参数path 为官方参数 ,参数saveDir为二维码图片存储路径 参数fileName 为二维码图片存储名称请勿包含名称
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (s *Server) GetQrcode(authorizerAccessToken string, path, saveDir, fileName *string) (filePath string, err error) {
	var (
		u        = WECHAT_API_URL + "/wxa/get_qrcode?"
		p        = core.AuthTokenUrlValues(authorizerAccessToken)
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

// GetQrcodeReturnBytes 获取体验版二维码 参数path 为官方参数 ,返回二进制byte数组,提供更灵活的使用
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (s *Server) GetQrcodeReturnBytes(authorizerAccessToken string, path *string) (qrcode []byte, err error) {
	var (
		u        = WECHAT_API_URL + "/wxa/get_qrcode?"
		p        = core.AuthTokenUrlValues(authorizerAccessToken)
		httpResp *http.Response
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
	qrcode, err = ioutil.ReadAll(httpResp.Body)
	return
}

type SubmitAuditReq struct {
	ItemList         []*SubmitAuditItem      `json:"item_list,omitempty"`
	PreviewInfo      *SubmitAuditPreviewInfo `json:"preview_info,omitempty"`
	VersionDesc      *string                 `json:"version_desc,omitempty"`
	FeedbackInfo     *string                 `json:"feedback_info,omitempty"`
	FeedbackStuff    *string                 `json:"feedback_stuff,omitempty"`
	UgcDeclare       *SubmitAuditUgcDeclare  `json:"ugc_declare,omitempty"`
	PrivacyApiNotUse *bool                   `json:"privacy_api_not_use,omitempty"`
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

// SubmitAudit 提交审核
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (s *Server) SubmitAudit(authorizerAccessToken string, req *SubmitAuditReq) (resp *SubmitAuditResp) {
	var (
		u = WECHAT_API_URL + "/wxa/submit_audit?"
	)
	resp = &SubmitAuditResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetAuditStatusResp struct {
	core.Error
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	Screenshot string `json:"screenshot"`
}

// GetAuditStatus 查询指定发布审核单的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html
func (s *Server) GetAuditStatus(authorizerAccessToken string, auditId int) (resp *GetAuditStatusResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/get_auditstatus?"
		req = &struct {
			AuditId int `json:"auditid"`
		}{AuditId: auditId}
	)
	resp = &GetAuditStatusResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetLatestAuditStatusResp struct {
	core.Error
	AuditId    int    `json:"auditid"`
	Status     int8   `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"screen_shot"`
}

// GetLatestAuditStatus 查询最新一次提交的审核状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_latest_auditstatus.html
func (s *Server) GetLatestAuditStatus(authorizerAccessToken string) (resp *GetLatestAuditStatusResp) {
	var (
		u = WECHAT_API_URL + "/wxa/get_latest_auditstatus?"
	)
	resp = &GetLatestAuditStatusResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// UndoCodeAudit 小程序审核撤回
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/undocodeaudit.html
func (s *Server) UndoCodeAudit(authorizerAccessToken string) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/undocodeaudit?"
	)
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// Release 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html
func (s *Server) Release(authorizerAccessToken string) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/release?"
		req = &struct{}{}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

// RevertCodeRelease 版本回退
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/revertCodeRelease.html
// Deprecated: 转用 RevertCodeReleaseV2 方法
func (s *Server) RevertCodeRelease(authorizerAccessToken string, appVersion int) (resp *core.Error) {
	var (
		u      = WECHAT_API_URL + "/wxa/revertcoderelease?"
		params = core.AuthTokenUrlValues(authorizerAccessToken)
	)
	//版本为零则回退到上一个版本
	if appVersion != 0 {
		params.Set("app_version", strconv.Itoa(appVersion))
	}
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, params, resp))
	return
}

type RevertCodeReleaseReq struct {
	Action     *string `json:"action,omitempty"`      //只能填get_history_version。表示获取可回退的小程序版本。该参数为 URL 参数，非 Body 参数。
	AppVersion *string `json:"app_version,omitempty"` //默认是回滚到上一个版本；也可回滚到指定的小程序版本，可通过get_history_version获取app_version。该参数为 URL 参数，非 Body 参数。
}

type RevertCodeReleaseResp struct {
	core.Error
	VersionList []*struct {
		AppVersion  int    `json:"app_version"`  //小程序版本
		UserVersion string `json:"user_version"` //模板版本号，开发者自定义字段
		UserDesc    string `json:"user_desc"`    //模板描述，开发者自定义字段
		CommitTime  int64  `json:"commit_time"`  //更新时间，时间戳
	} `json:"version_list,omitempty"`
}

// RevertCodeReleaseV2 版本回退
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/revertCodeRelease.html
func (s *Server) RevertCodeReleaseV2(authorizerAccessToken string, req *RevertCodeReleaseReq) (resp *RevertCodeReleaseResp) {
	var (
		u      = WECHAT_API_URL + "/wxa/revertcoderelease?"
		params = core.AuthTokenUrlValues(authorizerAccessToken)
	)
	//版本为零则回退到上一个版本
	if req != nil {
		if req.AppVersion != nil {
			params.Set("app_version", *req.AppVersion)
		}
		if req.Action != nil {
			params.Set("action", *req.Action)
		}
	}
	resp = &RevertCodeReleaseResp{}
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

// GetRevertCodeRelease 获取可回退的小程序版本
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_history_version.html
func (s *Server) GetRevertCodeRelease(authorizerAccessToken string) (resp *GetRevertCodeReleaseResp) {
	var (
		u      = WECHAT_API_URL + "/wxa/revertcoderelease?"
		params = core.AuthTokenUrlValues(authorizerAccessToken)
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

// GetPaidUnionId 支付后获取用户 Unionid 接口
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/User_Management.html
func (s *Server) GetPaidUnionId(authorizerAccessToken string, req *GetPaidUnionIdReq) (resp *GetPaidUnionIdResp) {
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

	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// GrayRelease 分阶段发布
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/grayrelease.html
func (s *Server) GrayRelease(authorizerAccessToken string, grayPercentage int8) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/grayrelease?"
		req = &struct {
			GrayPercentage int8 `json:"gray_percentage"`
		}{GrayPercentage: grayPercentage}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
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

// GetGrayReleasePlan 查询当前分阶段发布详情
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getgrayreleaseplan.html
func (s *Server) GetGrayReleasePlan(authorizerAccessToken string) (resp *GetGrayReleasePlanResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getgrayreleaseplan?"
	)
	resp = &GetGrayReleasePlanResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// RevertGrayRelease 取消分阶段发布
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/revertgrayrelease.html
func (s *Server) RevertGrayRelease(authorizerAccessToken string) (resp *core.Error) {
	var (
		u = WECHAT_API_URL + "/wxa/revertgrayrelease?"
	)
	resp = &core.Error{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// ChangeVisitStatus 修改小程序服务状态
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/change_visitstatus.html
func (s *Server) ChangeVisitStatus(authorizerAccessToken string, action string) (resp *core.Error) {
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
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
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

// GetWeappSupportVersion 查询当前设置的最低基础库版本及各版本用户占比
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/getweappsupportversion.html
func (s *Server) GetWeappSupportVersion(authorizerAccessToken string) (resp *GetWeappSupportVersionResp) {
	var (
		u   = CGIUrl + "/wxopen/getweappsupportversion?"
		req = &struct{}{}
	)

	resp = &GetWeappSupportVersionResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

// SetWeappSupportVersion 设置最低基础库版本
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/setweappsupportversion.html
func (s *Server) SetWeappSupportVersion(authorizerAccessToken string, version string) (resp *core.Error) {
	var (
		u   = CGIUrl + "/wxopen/setweappsupportversion?"
		req = &struct {
			Version string `json:"version"`
		}{Version: version}
	)

	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type QueryQuotaResp struct {
	core.Error
	Rest         int `json:"rest"`          //quota剩余值
	Limit        int `json:"limit"`         //当月分配quota
	SpeedupRest  int `json:"speedup_rest"`  //剩余加急次数
	SpeedupLimit int `json:"speedup_limit"` //当月分配加急次数
}

// QueryQuota 查询服务商的当月提审限额（quota）和加急次数
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/query_quota.html
func (s *Server) QueryQuota(authorizerAccessToken string) (resp *QueryQuotaResp) {
	var (
		u = WECHAT_API_URL + "/wxa/queryquota?"
	)
	resp = &QueryQuotaResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

// SpeedupAudit 加急审核申请
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/speedup_audit.html
func (s *Server) SpeedupAudit(authorizerAccessToken string, auditId int) (resp *core.Error) {
	var (
		u   = WECHAT_API_URL + "/wxa/speedupaudit?"
		req = &struct {
			AuditId int `json:"auditid"`
		}{AuditId: auditId}
	)

	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
