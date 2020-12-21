package wechat3rd

import "github.com/l306287405/wechat3rd/core"

// 授权方信息
const (
	AuthorizerInfoUrl      = wechatApiUrl + "/cgi-bin/component/api_get_authorizer_info?component_access_token=%s"
	AuthorizerOptionUrl    = wechatApiUrl + "/cgi-bin/component/api_get_authorizer_option?component_access_token=%s"
	SetAuthorizerOptionUrl = wechatApiUrl + "/cgi-bin/component/api_set_authorizer_option?component_access_token=%s"
	AuthorizerListUrl      = wechatApiUrl + "/cgi-bin/component/api_get_authorizer_list?component_access_token=%s"
)

type AuthorizerInfo struct {
}
type AuthorizerInfoRequest struct {
	ComponentAppid  string `json:"component_appid"`
	AuthorizerAppid string `json:"authorizer_appid"`
}

type AuthorizerInfoResponse struct {
	core.Error
	AuthorizerInfo struct {
		// 小程序独有
		Signature       string `json:"signature"`
		Miniprograminfo struct {
			Network struct {
				RequestDomain   []string `json:"RequestDomain"`
				WsRequestDomain []string `json:"WsRequestDomain"`
				UploadDomain    []string `json:"UploadDomain"`
				DownloadDomain  []string `json:"DownloadDomain"`
			} `json:"network"`
			Categories []struct {
				First  string `json:"first"`
				Second string `json:"second"`
			} `json:"categories"`
			VisitStatus int `json:"visit_status"`
		} `json:"miniprograminfo"`

		// 都存在的
		//昵称
		NickName        string `json:"nick_name"`
		HeadImg         string `json:"head_img"`
		//公众号类型  --公众号独有
		ServiceTypeInfo struct {
			ID int `json:"id"`
		} `json:"service_type_info"`
		// 认证类型
		VerifyTypeInfo struct {
			ID int `json:"id"`
		} `json:"verify_type_info"`
		//原始 ID
		UserName      string `json:"user_name"`
		// 主题名称
		PrincipalName string `json:"principal_name"`
		//用以了解功能的开通状况（0代表未开通，1代表已开通），详见business_info 说明
		BusinessInfo  struct {
			OpenStore int `json:"open_store"`
			OpenScan  int `json:"open_scan"`
			OpenPay   int `json:"open_pay"`
			OpenCard  int `json:"open_card"`
			OpenShake int `json:"open_shake"`
		} `json:"business_info"`
		Alias     string `json:"alias"`
		//二维码图片的 URL，开发者最好自行也进行保存
		QrcodeURL string `json:"qrcode_url"`
	} `json:"authorizer_info"`
	AuthorizationInfo struct {
		AuthorizationAppid string `json:"authorization_appid"`
		FuncInfo           []struct {
			FuncscopeCategory struct {
				ID int `json:"id"`
			} `json:"funcscope_category"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

// 获取授权法信息
func (srv *Server) AuthorizerInfo(authorizerAppid string) (*AuthorizerInfoResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &AuthorizerInfoRequest{
		ComponentAppid:  srv.cfg.AppID,
		AuthorizerAppid: authorizerAppid,
	}
	resp := &AuthorizerInfoResponse{}
	err = core.PostJson(getCompleteUrl(AuthorizerInfoUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type AuthorizeOption string

// option类型
//var (
//	AuthorizeOptionLocal           AuthorizeOption = "location_report"
//	AuthorizeOptionVoiceRecognize  AuthorizeOption = "voice_recognize"
//	AuthorizeOptionCustomerService AuthorizeOption = "customer_service"
//)

type AuthorizerOptionRequest struct {
	ComponentAppid  string          `json:"component_appid"`
	AuthorizerAppid string          `json:"authorizer_appid"`
	OptionName      AuthorizeOption `json:"option_name"`
}

type AuthorizerOptionResponse struct {
	core.Error
	AuthorizerAppid string `json:"authorizer_appid"`
	OptionName      string `json:"option_name"`
	OptionValue     string `json:"option_value"`
}

// 获取选项信息
func (srv *Server) AuthorizerOption(authorizerAppid string, optionName AuthorizeOption) (*AuthorizerOptionResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &AuthorizerOptionRequest{
		ComponentAppid:  srv.cfg.AppID,
		AuthorizerAppid: authorizerAppid,
		OptionName:      optionName,
	}
	resp := &AuthorizerOptionResponse{}
	err = core.PostJson(getCompleteUrl(AuthorizerOptionUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type SetAuthorizerOptionRequest struct {
	AuthorizerOptionRequest
	OptionValue string `json:"option_name"`
}

type SetAuthorizerOptionResponse struct {
	core.Error
}

// 设置选项信息
func (srv *Server) SetAuthorizerOption(authorizerAppid string, optionName AuthorizeOption, optionValue string) (*SetAuthorizerOptionResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &SetAuthorizerOptionRequest{
		AuthorizerOptionRequest: AuthorizerOptionRequest{
			ComponentAppid:  srv.cfg.AppID,
			AuthorizerAppid: authorizerAppid,
			OptionName:      optionName,
		},
		OptionValue: optionValue,
	}
	resp := &SetAuthorizerOptionResponse{}
	err = core.PostJson(getCompleteUrl(SetAuthorizerOptionUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type AuthorizerListRequest struct {
	ComponentAppid string `json:"component_appid"`
	Offset         int    `json:"offset"`
	Count          int    `json:"count"`
}

type AuthorizerListResponse struct {
	core.Error
	TotalCount int `json:"total_count"`
	List       []struct {
		AuthorizerAppid string `json:"authorizer_appid"`
		RefreshToken    string `json:"refresh_token"`
		AuthTime        int    `json:"auth_time"`
	} `json:"list"`
}

// 拉取用户授权列表
func (srv *Server) AuthorizerList(offset, count int) (*AuthorizerListResponse, error) {
	accessToken, err := srv.Token()
	if err != nil {
		return nil, err
	}
	req := &AuthorizerListRequest{
		ComponentAppid: srv.cfg.AppID,
		Offset:         offset,
		Count:          count,
	}
	resp := &AuthorizerListResponse{}
	err = core.PostJson(getCompleteUrl(AuthorizerListUrl, accessToken), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
