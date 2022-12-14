package wechat3rd

// https://github.com/fastwego/wxopen/blob/master/type/type_platform/type_platform.go

// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"encoding/xml"
)

const (
	EventTypeComponentVerifyTicket = "component_verify_ticket"
	EventTypeAuthorized            = "authorized"
	EventTypeUnauthorized          = "unauthorized"
	EventTypeUpdateAuthorized      = "updateauthorized"
	// 创建试用小程序成功/失败
	// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html
	EventTypeThirdFastRegisterBetaApp = "notify_third_fastregisterbetaapp"
	// 试用小程序快速认证
	// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastverify.html
	EventTypeThirdFastVerifyBetaApp = "notify_third_fastverifybetaapp"
	// 注册审核事件推送(快速创建个人小程序/快速注册企业小程序)
	// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html
	// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
	EventTypeThirdFastRegister = "notify_third_fasteregister"

	//代码审核结果推送
	//审核通过
	EventTypeWeappAuditSuccess = "weapp_audit_success"

	//审核不通过
	EventTypeWeappAuditFail = "weapp_audit_fail"

	//审核延后
	EventTypeWeappAuditDelay = "weapp_audit_delay"
)

type Event struct {
	XMLName    xml.Name `xml:"xml"`
	AppId      string
	CreateTime string
	InfoType   string
}

/*
<xml>
<AppId>some_appid</AppId>
<CreateTime>1413192605</CreateTime>
<InfoType>component_verify_ticket</InfoType>
<ComponentVerifyTicket>some_verify_ticket</ComponentVerifyTicket>
</xml>
*/
type EventComponentVerifyTicket struct {
	Event
	ComponentVerifyTicket string
}

/*
授权成功通知
<xml>

	<AppId>第三方平台appid</AppId>
	<CreateTime>1413192760</CreateTime>
	<InfoType>authorized</InfoType>
	<AuthorizerAppid>公众号appid</AuthorizerAppid>
	<AuthorizationCode>授权码</AuthorizationCode>
	<AuthorizationCodeExpiredTime>过期时间</AuthorizationCodeExpiredTime>
	<PreAuthCode>预授权码</PreAuthCode>

<xml>
*/
type EventAuthorized struct {
	Event
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

/*
取消授权通知
<xml>

	<AppId>第三方平台appid</AppId>
	<CreateTime>1413192760</CreateTime>
	<InfoType>unauthorized</InfoType>
	<AuthorizerAppid>公众号appid</AuthorizerAppid>

</xml>
*/
type EventUnauthorized struct {
	Event
	AuthorizerAppid string
}

/*
授权更新通知
<xml>

	<AppId>第三方平台appid</AppId>
	<CreateTime>1413192760</CreateTime>
	<InfoType>updateauthorized</InfoType>
	<AuthorizerAppid>公众号appid</AuthorizerAppid>
	<AuthorizationCode>授权码</AuthorizationCode>
	<AuthorizationCodeExpiredTime>过期时间</AuthorizationCodeExpiredTime>
	<PreAuthCode>预授权码</PreAuthCode>

<xml>
*/
type EventUpdateAuthorized struct {
	Event
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

/*
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastregister.html
<xml>

	    <AppId><![CDATA[第三方平台appid]]></AppId>
	    <CreateTime>1535442403</CreateTime>
	    <InfoType><![CDATA[notify_third_fastregisterbetaapp]]></InfoType>
	    <appid>创建小程序appid<appid>
	    <status>0</status>
	    <msg>OK</msg>
	    <info>
			<unique_id><![CDATA[unique_id]]></unique_id>
			<name><![CDATA[小程序名称]]></name>
	    </info>

</xml>
*/
type EventThirdFastRegisterBetaApp struct {
	Event
	MpAppid string `xml:"appid"`
	Status  int    `xml:"status"`
	Msg     string `xml:"msg"`
	Info    struct {
		UniqueID string `xml:"unique_id"`
		Name     string `xml:"name"`
	} `xml:"info"`
}

/*
*
https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/beta_Mini_Programs/fastverify.html
<xml>

	    <AppId><![CDATA[第三方平台appid]]></AppId>
	    <CreateTime>1535442403</CreateTime>
	    <InfoType><![CDATA[notify_third_fastverifybetaapp]]></InfoType>
	    <appid>小程序appid<appid>
	    <status>0</status>
	    <msg>OK</msg>
	    <info>
			<name><![CDATA[企业名称]]></name>
			<code><![CDATA[企业代码]]></code>
			<code_type>1</code_type>
			<legal_persona_wechat><![CDATA[法人微信号]]></legal_persona_wechat>
			<legal_persona_name><![CDATA[法人姓名]]></legal_persona_name>
			<component_phone><![CDATA[第三方联系电话]]></component_phone>
	    </info>

</xml>
*
*/
type EventThirdFastVerifyBetaApp struct {
	Event
	MpAppid string `xml:"appid"`
	Status  int    `xml:"status"`
	Msg     string `xml:"msg"`
	Info    struct {
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int    `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
	} `xml:"info"`
}

/*
*
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/fastregisterpersonalweapp.html
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Register_Mini_Programs/Fast_Registration_Interface_document.html
<xml>

	    <AppId><![CDATA[第三方平台appid]]></AppId>
	    <CreateTime>1535442403</CreateTime>
	    <InfoType><![CDATA[notify_third_fasteregister]]></InfoType>
	    <appid>创建小程序appid</appid>
	    <status>0</status>
	    <auth_code>xxxxx第三方授权码</auth_code>
	    <msg>OK</msg>
	    <info>
			<wxuser><![CDATA[用户微信号]]></wxuser>
			<idname><![CDATA[用户姓名]]></wxidnnn>
			<component_phone><![CDATA[第三方联系电话]]></component_phone>
	    </info>

</xml>
<xml>

	    <AppId><![CDATA[第三方平台appid]]></AppId>
	    <CreateTime>1535442403</CreateTime>
	    <InfoType><![CDATA[notify_third_fasteregister]]></InfoType>
	    <appid>创建小程序appid</appid>
	    <status>0</status>
	    <auth_code>xxxxx第三方授权码</auth_code>
	    <msg>OK</msg>
	    <info>
			<name><![CDATA[企业名称]]></name>
			<code><![CDATA[企业代码]]></code>
			<code_type>1</code_type>
			<legal_persona_wechat><![CDATA[法人微信号]]></legal_persona_wechat>
			<legal_persona_name><![CDATA[法人姓名]]></legal_persona_name>
			<component_phone><![CDATA[第三方联系电话]]></component_phone>
	    </info>

</xml>
*
*/
type EventThirdFastRegister struct {
	Event
	MpAppid  string `xml:"appid"`
	Status   int    `xml:"status"`
	Msg      string `xml:"msg"`
	AuthCode string `xml:"auth_code"`
	Info     struct {
		WxUser             string `xml:"wxuser"`
		IdName             string `xml:"idname"`
		Name               string `xml:"name"`
		Code               string `xml:"code"`
		CodeType           int    `xml:"code_type"`
		LegalPersonaWechat string `xml:"legal_persona_wechat"`
		LegalPersonaName   string `xml:"legal_persona_name"`
		ComponentPhone     string `xml:"component_phone"`
	} `xml:"info"`
}
