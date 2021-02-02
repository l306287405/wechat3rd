package wechat3rd

type MixedMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string   `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string   `xml:"MsgType"      json:"MsgType"`
	Event        string   `xml:"Event" json:"Event"`

	//echo
	EchoStr string `xml:"-" json:"-"`

	MsgId        int64   `xml:"MsgId"        json:"MsgId"`        // request
	Content      string  `xml:"Content"      json:"Content"`      // request
	MediaId      string  `xml:"MediaId"      json:"MediaId"`      // request
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`       // request
	Format       string  `xml:"Format"       json:"Format"`       // request
	Recognition  string  `xml:"Recognition"  json:"Recognition"`  // request
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"` // request
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`   // request
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`   // request
	Scale        int     `xml:"Scale"        json:"Scale"`        // request
	Label        string  `xml:"Label"        json:"Label"`        // request
	Title        string  `xml:"Title"        json:"Title"`        // request
	Description  string  `xml:"Description"  json:"Description"`  // request
	URL          string  `xml:"Url"          json:"Url"`          // request
	EventKey     string  `xml:"EventKey"     json:"EventKey"`     // request, menu
	Ticket       string  `xml:"Ticket"       json:"Ticket"`       // request
	Latitude     float64 `xml:"Latitude"     json:"Latitude"`     // request
	Longitude    float64 `xml:"Longitude"    json:"Longitude"`    // request
	Precision    float64 `xml:"Precision"    json:"Precision"`    // request
	BizMsgMenuId int64   `xml:"bizmsgmenuid" json:"bizmsgmenuid"` // request
	// menu
	MenuId       int64 `xml:"MenuId" json:"MenuId"`
	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo,omitempty" json:"ScanCodeInfo,omitempty"`
	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
	SendLocationInfo *struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		PoiName   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo,omitempty" json:"SendLocationInfo,omitempty"`

	MsgID  int64  `xml:"MsgID"  json:"MsgID"`  // template, mass
	Status string `xml:"Status" json:"Status"` // template, mass
	// shakearound
	ChosenBeacon *struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"ChosenBeacon,omitempty" json:"ChosenBeacon,omitempty"`
	AroundBeacons []struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`

	UnionId string `xml:"UnionId"              json:"UnionId"` // unionId

	// openapi 推送
	AppId string `xml:"AppId" json:"AppId"`
	//CreateTime int32 `xml:"CreateTime" json:"CreateTime"`
	InfoType              string `xml:"InfoType" json:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"`

	AuthorizerAppid              string `json:"AuthorizerAppid" xml:"AuthorizerAppid"`
	AuthorizationCode            string `json:"AuthorizationCode" xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime string `json:"AuthorizationCodeExpiredTime" xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string `json:"PreAuthCode" xml:"PreAuthCode"`
}
