package wechat3rd

import "encoding/xml"

/**
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/audit_event.html
<xml>
  <ToUserName><![CDATA[gh_fb9688c2a4b2]]></ToUserName>
  <FromUserName><![CDATA[od1P50M-fNQI5Gcq-trm4a7apsU8]]></FromUserName>
  <CreateTime>1488856741</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[weapp_audit_success]]></Event>
  <SuccTime>1488856741</SuccTime>
</xml>
**/

// EventWeappAuditSuccess 代码审核结果推送
// 审核通过
type EventWeappAuditSuccess struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   //小程序的原始 ID
	FromUserName string   `xml:"FromUserName"` //发送方帐号（一个 OpenID，此时发送方是系统帐号）
	CreateTime   int64    `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`      //消息类型，event
	Event        string   `xml:"Event"`        //事件类型，weapp_audit_success
	SuccTime     int64    `xml:"SuccTime"`     //审核成功时间，单位：秒
}

// EventWeappAuditFail 代码审核结果推送
// 审核不通过
type EventWeappAuditFail struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   //小程序的原始 ID
	FromUserName string   `xml:"FromUserName"` //发送方帐号（一个 OpenID，此时发送方是系统帐号）
	CreateTime   int64    `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`      //消息类型，event
	Event        string   `xml:"Event"`        //事件类型，weapp_audit_success
	Reason       string   `xml:"Reason"`       //审核失败原因
	FailTime     int64    `xml:"FailTime"`     //审核失败时间，单位：秒
	ScreenShot   string   `xml:"ScreenShot"`   //审核不通过的截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

// EventWeappAuditDelay 代码审核结果推送
// 审核延后
type EventWeappAuditDelay struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   //小程序的原始 ID
	FromUserName string   `xml:"FromUserName"` //发送方帐号（一个 OpenID，此时发送方是系统帐号）
	CreateTime   int64    `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`      //消息类型，event
	Event        string   `xml:"Event"`        //事件类型，weapp_audit_success
	Reason       string   `xml:"Reason"`       //审核失败原因
	DelayTime    int64    `xml:"DelayTime"`    //审核延后时的时间戳
}
