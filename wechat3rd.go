package wechat3rd

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/l306287405/wechat3rd/core"
	"github.com/l306287405/wechat3rd/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// open api 配置
type Config struct {
	AppID     string
	AppSecret string
	AESKey    string
	Token     string
	//RedirectUrl    string
}

func (c *Config) check() error{
	if len(c.AESKey) != 43 {
		//log.Fatalln("the length of base64AESKey must equal to 43")
		return errors.New("the length of base64AESKey must equal to 43")
	}

	if len(c.Token) != 32{
		return errors.New("token was not set for Server, see NewServer function or Server.SetToken method")
	}

	if c.AppID == ""{
		return errors.New("appid was not set for Server")
	}

	if c.AppSecret == ""{
		return errors.New("app secret was not set for Server!")
	}
	return nil
}

//type Handler func(c *MixedMsg)

type Server struct {
	sync.Mutex
	cfg          Config
	//handlerMap   map[string]Handler //方法处理
	DecodeAesKey []byte
	errorHandler WechatErrorer           // 错误处理
	TicketServer // ticket存储
	// 获取token
	AccessTokenServer

}

const (
	// InfoTypeVerifyTicket 返回ticket
	InfoTypeVerifyTicket = "component_verify_ticket"
	// InfoTypeAuthorized 授权
	InfoTypeAuthorized = "authorized"
	// InfoTypeUnauthorized 取消授权
	InfoTypeUnauthorized = "unauthorized"
	// InfoTypeUpdateAuthorized 更新授权
	InfoTypeUpdateAuthorized = "updateauthorized"

	wechatApiUrl = "https://api.weixin.qq.com"
)

func (srv *Server) getAESKey() []byte {
	return srv.DecodeAesKey
}
func (srv *Server) getToken() string {
	return srv.cfg.Token
}

type cipherRequestHttpBody struct {
	XMLName            struct{} `xml:"xml"`
	ToUserName         string   `xml:"ToUserName"`
	AppId              string   `xml:"AppId"` // openapi use
	Base64EncryptedMsg []byte   `xml:"Encrypt"`
}

func NewService(cfg Config, ticket TicketServer, tokenService AccessTokenServer, errHandler WechatErrorer) (s *Server,err error) {
	err=cfg.check()
	if err!=nil{
		return nil, err
	}

	if errHandler == nil {
		errHandler = DefaultErrorHandler
	}
	if ticket == nil {
		ticket = defaultTicketServerHander
	}
	if tokenService == nil {
		tokenService = &DefaultAccessTokenServer{TicketServer: ticket, AppID: cfg.AppID, AppSecret: cfg.AppSecret}
	}
	s = &Server{
		cfg:               cfg,
		errorHandler:      errHandler,
		//handlerMap:        make(map[string]Handler),
		TicketServer:      ticket,
		AccessTokenServer: tokenService,
	}
	s.DecodeAesKey,err=base64.StdEncoding.DecodeString(s.cfg.AESKey + "=")
	if err != nil {
		return nil, errors.New("Decode base64AESKey failed: "+err.Error())
	}

	return s,nil
}

//func (srv *Server) AddHander(t string, hander Handler) {
//	srv.handlerMap[t] = hander
//}

func (srv *Server) ServeHTTP(r *http.Request) (resp *MixedMsg,err error) {
	var(
		query=r.URL.Query()

		wantSignature string
		haveSignature = query.Get("signature")
		timestamp = query.Get("timestamp")
		nonce = query.Get("nonce")

		//get
		echostr = query.Get("echostr")

		//post
		wantMsgSignature string
		haveMsgSignature = query.Get("msg_signature")
		encryptType = query.Get("encrypt_type")

		//handle vars
		data []byte
		requestHttpBody = &cipherRequestHttpBody{}
		encryptedMsg []byte
		encryptedMsgLen int
		msgPlaintext, haveAppIdBytes []byte
		//hand Handler
		//exist bool
	)

	if haveSignature == "" {
		err=errors.New("not found signature query parameter")
		return
	}
	if timestamp == "" {
		err=errors.New("not found timestamp query parameter")
		return
	}
	if nonce == "" {
		err=errors.New("not found nonce query parameter")
		return
	}

	wantSignature = util.Sign(srv.getToken(), timestamp, nonce)
	if haveSignature != wantSignature {
		return nil,errors.New("sign error")
	}

	//如果是验证url有效性 则echo即可
	if r.Method == "GET"{
		if echostr == "" {
			err=errors.New("not found echostr query parameter")
			return
		}
		resp = &MixedMsg{EchoStr: echostr}
		return
	}

	//进入事件执行
	if encryptType!="aes"{
		err=errors.New("unknown encrypt_type: "+encryptType)
		return
	}
	if haveMsgSignature == "" {
		err=errors.New("not found msg_signature query parameter")
		return
	}

	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, requestHttpBody)
	if err != nil {
		return
	}

	wantMsgSignature = util.MsgSign(srv.getToken(), timestamp, nonce, string(requestHttpBody.Base64EncryptedMsg))
	if haveMsgSignature!=wantMsgSignature {
		err= errors.New("check msg_signature failed, have: "+haveMsgSignature+", want: "+wantMsgSignature)
		return
	}

	encryptedMsg = make([]byte, base64.StdEncoding.DecodedLen(len(requestHttpBody.Base64EncryptedMsg)))
	encryptedMsgLen, err = base64.StdEncoding.Decode(encryptedMsg, requestHttpBody.Base64EncryptedMsg)
	if err != nil {
		return
	}
	encryptedMsg = encryptedMsg[:encryptedMsgLen]

	_, msgPlaintext, haveAppIdBytes, err = util.AESDecryptMsg(encryptedMsg, srv.getAESKey())
	if err != nil {
		return
	}

	if string(haveAppIdBytes)!=srv.cfg.AppID {
		err = errors.New("the message AppId mismatch, have: "+string(haveAppIdBytes)+", want: "+srv.cfg.AppID)
		return
	}
	resp=&MixedMsg{}
	if err = xml.Unmarshal(msgPlaintext, resp); err != nil {
		return
	}
	//
	//hand, exist = srv.handlerMap[resp.InfoType]
	//if !exist {
	//	err = errors.New("match handler failed :"+resp.InfoType)
	//	return
	//}
	//hand(resp)
	return
}

//用于解密数据
func (s *Server) AESDecryptData(cipherText , iv []byte)(rawData []byte,err error){
	return util.AESDecryptData(cipherText,s.getAESKey(),iv)
}

type Jscode2SessionResp struct {
	Openid string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Errcode string `json:"errcode,omitempty"`
	Errmsg string `json:"errmsg,omitempty"`
}

func (s *Server) Jscode2Session(appId ,jsCode string) (resp *Jscode2SessionResp,err error){
	var(
		req = make(url.Values)
		u = "https://api.weixin.qq.com/sns/component/jscode2session?"
	)
	resp = &Jscode2SessionResp{}
	req.Set("appid",appId)
	req.Set("js_code",jsCode)
	req.Set("grant_type","authorization_code")
	req.Set("component_appid",s.cfg.AppID)
	req.Set("component_access_token",s.getToken())

	err=core.GetRequest(u,req,resp)
	return
}