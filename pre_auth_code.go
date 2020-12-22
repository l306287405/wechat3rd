package wechat3rd

import (
	"fmt"
	"github.com/l306287405/wechat3rd/core"
	"log"
)

type AuthType string

const (
	PreAuthCodeUrl  = wechatApiUrl + "/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	AuthPageUrl     = wechatApiUrl + "/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s"
	QueryAuthUrl    = wechatApiUrl + "/cgi-bin/component/api_query_auth?component_access_token=%s"
	RefreshTokenUrl = wechatApiUrl + "/cgi-bin/component/api_authorizer_token?component_access_token=%s"

	PreAuthAuthTypeAll     AuthType = "3" // 全部
	PreAuthAuthTypeMinapp  AuthType = "2" // 小程序
	PreAuthAuthTypeService AuthType = "1" // 公众号
)

type PreAuthCodeRequest struct {
	ComponentAppid string `json:"component_appid"`
}

type PreAuthCodeResponse struct {
	core.Error
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

func (srv *Server) PreAuthCode() (*PreAuthCodeResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &PreAuthCodeRequest{
		ComponentAppid: srv.cfg.AppID,
	}
	log.Println("token:",accessToken)
	resp := &PreAuthCodeResponse{}
	err = core.PostJson(getCompleteUrl(PreAuthCodeUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *Server) AuthUrl(redirectUri string, authType AuthType) string {
	pcode, err := srv.PreAuthCode()
	if err != nil {
		return ""
	}
	return fmt.Sprintf(AuthPageUrl, srv.cfg.AppID, pcode.PreAuthCode, redirectUri, authType)
}

type QueryAuthRequest struct {
	ComponentAppid    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}
type QueryAuthResponse struct {
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
func (srv *Server) QueryAuth(code string) (*QueryAuthResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &QueryAuthRequest{
		ComponentAppid:    srv.cfg.AppID,
		AuthorizationCode: code,
	}
	resp := &QueryAuthResponse{}
	err = core.PostJson(getCompleteUrl(QueryAuthUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type RefreshTokenRequest struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
type RefreshTokenResponse struct {
	core.Error
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// 刷新token
func (srv *Server) RefreshToken(appID, refreshToken string) (*RefreshTokenResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &RefreshTokenRequest{
		ComponentAppid:         srv.cfg.AppID,
		AuthorizerAppid:        appID,
		AuthorizerRefreshToken: refreshToken,
	}
	resp := &RefreshTokenResponse{}
	err = core.PostJson(getCompleteUrl(RefreshTokenUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getCompleteUrl(uri, token string) string {
	return fmt.Sprintf(uri, token)
}
