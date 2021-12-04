package wechat3rd

import (
	"github.com/l306287405/wechat3rd/core"
	"net/url"
)

type MediaUploadReq struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type MediaUploadResp struct {
	core.Error
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

//新增临时素材
//https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (s *Server) MediaUpload(accessToken string, req *MediaUploadReq) (resp *MediaUploadResp) {
	var (
		u = CGIUrl + "/media/upload?"
	)
	u = s.AuthToken2url(u, accessToken) + "&type=" + req.Type
	resp = &MediaUploadResp{}
	resp.Err(core.PostFile(u, req.Path, "media", resp))
	return
}

type MediaGetResp struct {
	core.Error
	VideoUrl string `json:"video_url"`
}

//获取临时素材
//https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html
func (s *Server) MediaGet(accessToken string, mediaId string) (resp *MediaGetResp) {
	var (
		u   = CGIUrl + "/media/get?"
		req = url.Values{}
	)
	req.Set("media_id", mediaId)
	resp = &MediaGetResp{}
	resp.Err(core.GetRequest(s.AuthToken2url(u, accessToken), req, resp))
	return
}

type MaterialItem struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumbMediaId"`
	ShowCoverPic     int8   `json:"showCoverPic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"contentSourceUrl"`
}

type GetMaterialResp struct {
	core.Error
	//图文内容响应结果
	NewsItem []*MaterialItem `json:"newsItem,omitempty"`

	//视频消息响应结果
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	DownUrl     *string `json:"downUrl,omitempty"`
}

//获取永久素材
//https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (s *Server) GetMaterial(accessToken string, mediaId string) (resp *GetMaterialResp) {
	var (
		u   = CGIUrl + "/material/get_material?"
		req = &struct {
			MediaId string `json:"mediaId"`
		}{MediaId: mediaId}
	)
	resp = &GetMaterialResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, accessToken), req, resp))
	return
}
