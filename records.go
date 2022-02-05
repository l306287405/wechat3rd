package wechat3rd

import "github.com/l306287405/wechat3rd/core"

type GetIllegalRecordsReq struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

type GetIllegalRecordsResp struct {
	core.Error
	Records []*IllegalRecord `json:"records"`
}

type IllegalRecord struct {
	IllegalRecordId string `json:"illegal_record_id"` //违规处罚记录id
	CreateTime      int64  `json:"create_time"`       //违规处罚时间
	IllegalReason   string `json:"illegal_reason"`    //违规原因
	IllegalContent  string `json:"illegal_content"`   //违规内容
	RuleUrl         string `json:"rule_url"`          //规则链接
	RuleName        string `json:"rule_name"`         //违反的规则名称
}

//获取小程序违规处罚记录
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/records/getillegalrecords.html
func (s *Server) GetIllegalRecords(authorizerAccessToken string, req *GetIllegalRecordsReq) (resp *GetIllegalRecordsResp) {
	var (
		u = WECHAT_API_URL + "/wxa/getillegalrecords?"
	)
	if req == nil {
		req = &GetIllegalRecordsReq{}
	}
	resp = &GetIllegalRecordsResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}

type GetAppealRecordsResp struct {
	core.Error
	Records []*AppealRecord `json:"records"`
}

type AppealRecord struct {
	AppealRecordId    int         `json:"appeal_record_id"`   //申诉单id
	AppealTime        int         `json:"appeal_time"`        //申诉时间
	AppealCount       int         `json:"appeal_count"`       //申诉次数
	AppealFrom        int         `json:"appeal_from"`        //申诉来源（0--用户，1--服务商）
	AppealStatus      int         `json:"appeal_status"`      //申诉状态
	AuditTime         int         `json:"audit_time"`         //审核时间
	AuditReason       string      `json:"audit_reason"`       //审核结果理由
	PunishDescription string      `json:"punish_description"` //处罚原因描述
	Materials         []*Material `json:"materials"`          //违规材料和申诉材料
}

type Material struct {
	IllegalMaterial struct {
		Content    string `json:"content"`
		ContentUrl string `json:"content_url"`
	} `json:"illegal_material"`
	AppealMaterial struct {
		Reason           string   `json:"reason"`
		ProofMaterialIds []string `json:"proof_material_ids"`
	} `json:"appeal_material"`
}

//获取小程序申诉记录
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/records/getappealrecords.html
func (s *Server) GetAppealRecords(authorizerAccessToken string, illegalRecordId string) (resp *GetAppealRecordsResp) {
	var (
		u   = WECHAT_API_URL + "/wxa/getappealrecords?"
		req = &struct {
			IllegalRecordId string `json:"illegal_record_id"`
		}{IllegalRecordId: illegalRecordId}
	)
	resp = &GetAppealRecordsResp{}
	resp.Err(core.PostJson(s.AuthToken2url(u, authorizerAccessToken), req, resp))
	return
}
