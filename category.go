package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type GetAllCategoryResp struct {
	core.Error
	CategoriesList struct {
		Categories []*AllCategory `json:"categories"`
	} `json:"categories_list"` //类目信息列表
}

type AllCategory struct {
	Id            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Level         string `json:"level,omitempty"`
	Father        int    `json:"father,omitempty"`
	Children      []int  `json:"children,omitempty"`
	SensitiveType int8   `json:"sensitive_type,omitempty"`
	Qualify       struct {
		ExterList []*struct {
			InnerList []*struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"inner_list"`
		} `json:"exter_list"`
	} `json:"qualify,omitempty"`
}

// 获取可以设置的所有类目
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getallcategories.html
func (s *Server) GetMiniProgramAllCategory(authorizerAccessToken string) (resp *GetAllCategoryResp) {
	var (
		u = CGIUrl + "/wxopen/getallcategories?"
	)
	resp = &GetAllCategoryResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

type GetCategoryResp struct {
	core.Error
	Categories []*struct {
		First       int    `json:"first"`
		FirstName   string `json:"first_name"`
		Second      int    `json:"second"`
		SecondName  string `json:"second_name"`
		AuditStatus int    `json:"audit_status"`
		AuditReason string `json:"audit_reason"`
	}
	Limit         int `json:"limit"`
	Quota         int `json:"quota"`
	CategoryLimit int `json:"category_limit"`
}

// 获取已设置的所有类目
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategory.html
func (s *Server) GetMiniProgramCategory(authorizerAccessToken string) (resp *GetCategoryResp) {
	var (
		u = CGIUrl + "/wxopen/getcategory?"
	)
	resp = &GetCategoryResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}

type GetCategoriesByTypeResp struct {
	core.Error
	CategoriesList struct {
		Categories []*Category `json:"categories"`
	} `json:"categories_list"` //类目信息列表
}

// 获取不同主体类型的类目
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/getcategorybytype.html
func (s *Server) GetMiniProgramCategoriesByType(authorizerAccessToken string, verifyType ...int8) (resp *GetCategoriesByTypeResp) {
	var (
		u   = CGIUrl + "/wxopen/getcategoriesbytype?"
		req = &struct {
			VerifyType int8 `json:"verify_type"`
		}{}
	)
	resp = &GetCategoriesByTypeResp{}
	if verifyType != nil {
		req.VerifyType = verifyType[0]
	}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type Categories struct {
	First      int                 `json:"first"`
	Second     int                 `json:"second"`
	Certicates []*CategoryCertCate `json:"certicates"`
}

type CategoryCertCate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 添加类目
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/addcategory.html
func (s *Server) AddMiniProgramCategory(authorizerAccessToken string, categories []*Categories) (resp *core.Error) {
	var (
		u   = CGIUrl + "/wxopen/addcategory?"
		req = &struct {
			Categories []*Categories `json:"categories"`
		}{Categories: categories}
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type DeleteCategoryReq struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

// 删除类目
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/deletecategory.html
func (s *Server) DeleteMiniProgramCategory(authorizerAccessToken string, req *DeleteCategoryReq) (resp *core.Error) {
	var (
		u = CGIUrl + "/wxopen/addcategory?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

// 修改类目资质信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/modifycategory.html
func (s *Server) ModifyMiniProgramCategory(authorizerAccessToken string, categories *Categories) (resp *core.Error) {
	var (
		u = CGIUrl + "/wxopen/modifycategory?"
	)
	resp = &core.Error{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), categories, resp))
	return
}

type CategoriesResp struct {
	core.Error
	CategoryList []*struct {
		FirstClass  string  `json:"first_class"`
		FirstId     int     `json:"first_id"`
		SecondClass string  `json:"second_class"`
		SecondId    int     `json:"second_id"`
		ThirdClass  *string `json:"third_class,omitempty"`
		ThirdId     *int    `json:"third_id,omitempty"`
	} `json:"category_list"`
}

// 获取审核时可填写的类目信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html
func (s *Server) GetMiniProgramAuditCategory(authorizerAccessToken string) (resp *CategoriesResp) {
	var (
		u = WECHAT_API_URL + "/wxa/get_category?"
	)
	resp = &CategoriesResp{}
	resp.Err(core.GetRequest(u, core.AuthTokenUrlValues(authorizerAccessToken), resp))
	return
}
