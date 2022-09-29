package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"io"
)

type SetPrivacySettingReq struct {
	PrivacyVer         *int                  `json:"privacy_ver,omitempty"`
	OwnerSetting       PrivacyOwnerSetting   `json:"owner_setting"`
	SettingList        []*PrivacySetting     `json:"setting_list,omitempty"`
	SdkPrivacyInfoList []*SdkPrivacyInfoList `json:"sdk_privacy_info_list,omitempty"`
}

type PrivacyOwnerSetting struct {
	ContactEmail         *string `json:"contact_email,omitempty"`
	ContactPhone         *string `json:"contact_phone,omitempty"`
	ContactQQ            *string `json:"contact_qq,omitempty"`
	ContactWeixin        *string `json:"contact_weixin,omitempty"`
	ExtFileMediaId       *string `json:"ext_file_media_id,omitempty"`
	NoticeMethod         string  `json:"notice_method"`
	StoreExpireTimestamp *string `json:"store_expire_timestamp,omitempty"`
}

type PrivacySetting struct {
	PrivacyKey   string  `json:"privacy_key"`
	PrivacyText  string  `json:"privacy_text"`
	PrivacyLabel *string `json:"privacy_label,omitempty"`
	PrivacyDesc  *string `json:"privacy_desc,omitempty"`
}

type SdkPrivacyInfoList struct {
	SdkName    string     `json:"sdk_name"`
	SdkBizName string     `json:"sdk_biz_name"`
	SdkList    []*SdkList `json:"sdk_list"`
}
type SdkList struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyText string `json:"privacy_text"`
}

//设置小程序用户隐私保护指引
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func (s *Server) SetPrivacySetting(authorizerAccessToken string, req *SetPrivacySettingReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/component/setprivacysetting?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetPrivacySettingResp struct {
	core.Error
	CodeExist    int                 `json:"code_exist"`
	PrivacyList  []string            `json:"privacy_list"`
	SettingList  []*PrivacySetting   `json:"setting_list"`
	UpdateTime   int64               `json:"update_time"`
	OwnerSetting PrivacyOwnerSetting `json:"owner_setting"`
	PrivacyDesc  struct {
		PrivacyDescList []*PrivacySetting `json:"privacy_desc_list"`
	} `json:"privacy_desc"`
	SdkPrivacyInfoList []*SdkPrivacyInfoList `json:"sdk_privacy_info_list"`
}

//查询小程序用户隐私保护指引
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/get_privacy_setting.html
func (s *Server) GetPrivacySetting(authorizerAccessToken string, privacyVer ...int) (resp *GetPrivacySettingResp) {
	var (
		u   = CGIUrl + "/component/getprivacysetting?"
		req = &struct {
			PrivacyVer *int `json:"privacy_ver,omitempty"`
		}{}
	)
	resp = &GetPrivacySettingResp{}
	if privacyVer != nil {
		req.PrivacyVer = &privacyVer[0]
	}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type UploadPrivacyExtFileResp struct {
	core.Error
	ExtFileMediaId string `json:"ext_file_media_id"`
}

//上传小程序用户隐私保护指引
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func (s *Server) UploadPrivacyExtFile(authorizerAccessToken string, filePath string) (resp *UploadPrivacyExtFileResp) {
	var (
		u = CGIUrl + "/component/uploadprivacyextfile?"
	)
	resp = &UploadPrivacyExtFileResp{}
	resp.Err(core.PostFile(s.AuthToken2url(u, authorizerAccessToken), filePath, "file", resp))
	return
}

//上传小程序用户隐私保护指引,直接传入文件io数据,文件获取交给上层业务代码,使用更灵活，更直接
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/privacy_config/set_privacy_setting.html
func (s *Server) UploadPrivacyExtFileDirectly(authorizerAccessToken string, file io.Reader, fileName string) (resp *UploadPrivacyExtFileResp) {
	var (
		u = CGIUrl + "/component/uploadprivacyextfile?"
	)
	resp = &UploadPrivacyExtFileResp{}
	resp.Err(core.PostFileDirectly(s.AuthToken2url(u, authorizerAccessToken), file, fileName, "file", resp))
	return
}
