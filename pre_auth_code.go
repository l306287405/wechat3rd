package wechat3rd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/l306287405/wechat3rd/core"
)

type AuthType string

const (
	PREAUTH_CODE_URL  = WECHAT_API_URL + "/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	WEB_AUTH_URL      = WECHAT_MP_URL + "/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s"
	MOBILE_AUTH_URL   = WECHAT_MP_URL + "/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s"
	QUERY_AUTH_URL    = WECHAT_API_URL + "/cgi-bin/component/api_query_auth?component_access_token=%s"
	REFRESH_TOKEN_URL = WECHAT_API_URL + "/cgi-bin/component/api_authorizer_token?component_access_token=%s"

	PREAUTH_AUTH_TYPE_All     AuthType = "3" // 全部
	PREAUTH_AUTH_TYPE_MINIAPP AuthType = "2" // 小程序
	PREAUTH_AUTH_TYPE_Service AuthType = "1" // 公众号
)

type PreAuthCodeReq struct {
	ComponentAppid string `json:"component_appid"`
}

type PreAuthCodeResp struct {
	core.Error
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

func (srv *Server) PreAuthCode() (resp *PreAuthCodeResp) {
	resp = &PreAuthCodeResp{}
	accessToken, err := srv.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	req := &PreAuthCodeReq{
		ComponentAppid: srv.cfg.AppID,
	}
	resp.Err(core.PostJson(getCompleteUrl(PREAUTH_CODE_URL, accessToken), req, resp))
	return
}

//说明
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Authorization_Process_Technical_Description.html
func (srv *Server) AuthUrl(isWebAuth bool, redirectUri string, authType AuthType, bizAppid *string) (u string, err error) {
	var (
		resp *PreAuthCodeResp
	)
	resp = srv.PreAuthCode()
	if !resp.Success() {
		err = errors.New(resp.ErrMsg)
		return
	}
	tPreAuthCode := url.QueryEscape(resp.PreAuthCode)
	redirectUri = url.QueryEscape(redirectUri)
	if isWebAuth {
		u = fmt.Sprintf(WEB_AUTH_URL, srv.cfg.AppID, tPreAuthCode, redirectUri, authType)
	} else {
		u = fmt.Sprintf(MOBILE_AUTH_URL, srv.cfg.AppID, tPreAuthCode, redirectUri, authType)
		if bizAppid != nil && *bizAppid != "" {
			u = u + "&biz_appid=" + *bizAppid
		}
		u += "#wechat_redirect"
	}

	return u, nil
}

type QueryAuthReq struct {
	ComponentAppid    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}
type QueryAuthResp struct {
	core.Error
	AuthorizationInfo struct {
		AuthorizerAppid        string `json:"authorizer_appid"`
		AuthorizerAccessToken  string `json:"authorizer_access_token"`
		ExpiresIn              int    `json:"expires_in"`
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
		FuncInfo               []struct {
			FuncscopeCategory struct {
				ID int `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// 返回授权数据
func (srv *Server) QueryAuth(code string) (resp *QueryAuthResp) {
	resp = &QueryAuthResp{}
	accessToken, err := srv.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	req := &QueryAuthReq{
		ComponentAppid:    srv.cfg.AppID,
		AuthorizationCode: code,
	}
	resp.Err(core.PostJson(getCompleteUrl(QUERY_AUTH_URL, accessToken), req, resp))
	return
}

type RefreshTokenReq struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
type RefreshTokenResp struct {
	core.Error
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// 刷新token
func (srv *Server) RefreshToken(appID, refreshToken string) (resp *RefreshTokenResp) {
	resp = &RefreshTokenResp{}
	accessToken, err := srv.Token()
	if err != nil {
		resp.Err(err)
		return
	}
	req := &RefreshTokenReq{
		ComponentAppid:         srv.cfg.AppID,
		AuthorizerAppid:        appID,
		AuthorizerRefreshToken: refreshToken,
	}
	resp.Err(core.PostJson(getCompleteUrl(REFRESH_TOKEN_URL, accessToken), req, resp))
	return
}

func getCompleteUrl(uri, token string) string {
	return fmt.Sprintf(uri, token)
}
