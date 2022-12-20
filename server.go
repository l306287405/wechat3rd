package wechat3rd

import (
	"encoding/xml"
)

// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorize_event.html
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket.html
/*
启用 加密模式 后 收到的 消息格式
<xml>
    <ToUserName><![CDATA[]]></ToUserName>
    <Encrypt><![CDATA[]]></Encrypt>
</xml>
*/
type EncryptMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

// ParseXML 解析微信推送过来的消息/事件
func ParseXML(body []byte) (m interface{}, err error) {
	event := &Event{}
	if err = xml.Unmarshal(body, event); err != nil {
		return
	}

	switch event.InfoType {
	case EventTypeComponentVerifyTicket:
		msg := &EventComponentVerifyTicket{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeAuthorized:
		msg := &EventAuthorized{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeUnauthorized:
		msg := &EventUnauthorized{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeUpdateAuthorized:
		msg := &EventUpdateAuthorized{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeThirdFastRegisterBetaApp:
		msg := &EventThirdFastRegisterBetaApp{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeThirdFastVerifyBetaApp:
		msg := &EventThirdFastVerifyBetaApp{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeThirdFastRegister:
		msg := &EventThirdFastRegister{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeWeappAuditSuccess:
		msg := &EventWeappAuditSuccess{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeWeappAuditFail:
		msg := &EventWeappAuditFail{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	case EventTypeWeappAuditDelay:
		msg := &EventWeappAuditDelay{}
		if err = xml.Unmarshal(body, msg); err != nil {
			return
		}
		return msg, nil
	}
	return
}
