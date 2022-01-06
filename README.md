### 项目地址
[l306287405/wechat3rd](https://github.com/l306287405/wechat3rd)

### 感谢
项目基于 [owen-gxz/open-wechat](https://github.com/owen-gxz/open-wechat) 做了原方法改动以及添加对新接口的支持并补充使用说明.

### 1.简介

微信公众平台-第三方平台（简称第三方平台）开放给所有通过开发者资质认证后的开发者使用。  
在得到公众号或小程序运营者（简称运营者）授权后，第三方平台开发者可以通过调用微信开放平台的接口能力，为公众号或小程序的运营者提供账号申请、小程序创建、技术开发、行业方案、活动营销、插件能力等全方位服务。  
同一个账号的运营者可以选择多家适合自己的第三方为其提供产品能力或委托运营。

### 2.使用引导

首先请认真阅读 [官方文档](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/Third_party_platform_appid.html)  
有任何问题请提issue, 我会尽快解决.

#### 2.1: 引入
    go get -u github.com/l306287405/wechat3rd@master
    or
    go get -u github.com/l306287405/wechat3rd@v1.7.2 (请选择最新版本)
    v1.6.0版本开始Service提供的所有方法仅在resp返回对象中提供错误信息

#### 2.2: 使用NewService方法来创建一个service

    NewService的参数分别是:
    Config: 配置信息
    TicketServer: 保存微信传输的ticket信息接口
    AccessTokenServer: 获取第三方平台的token接口
    WechatErrorer: 错误信息的处理

    // 除Config外的其它参数传nil则使用默认配置.  该处代码你应该使用单例模式或服务池方式来管理
    service,err:=wechat3rd.NewService(wechat3rd.Config{
        AppID:     os.Getenv("WX_OPEN_APP_ID"), //第三方平台appid
        AppSecret: os.Getenv("WX_OPEN_APP_SECRET"), //第三方平台app_secret
        AESKey:    os.Getenv("WX_AES_KEY"), //第三方平台消息加解密Key
        Token:     os.Getenv("WX_AES_TOKEN"), //消息校验Token
    },nil,nil,nil)

    if err!=nil{
        panic("wechat3rd初始化失败:"+err.Error())
    }

    // 设置第三方平台的ticket 注意该处为示例代码 请替换缓存为自己的缓存服务.
    // https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket.html
    // 解释: 当第三方平台应用被创建时,微信每隔10分钟会向用户填写的 授权事件接收URL 发起POST请求,
    //      请求中带有12小时有效期的 component_verify_ticket 是服务使用加解密的必备参数.
    //      当服务重启时该参数会丢失从而导致服务请求失败,被动的方式是等下一次微信发起授权事件的请求并设置ticket,
    //      主动的方式则是将微信发起的ticket缓存起来并在服务启动时获取并设置.
    ticket:=cache.Get("cachekey_of_wechat3rd_ticket").Val()
    if ticket!=""{
        err=service.SetTicket(ticket)
        if err!=nil{
            common.Log.Error("设置微信isp的ticket失败:",err.Error()," 等待自动获取")
        }
    }

#### 2.3: 使用service来接收并验证票据

    // 使用你的Golang框架方法获取 *http.Request 这里用iris演示
    var r = c.Ctx.Request()
    
    // 该处service应该是从你的单例方法或者服务池中获取
    resp,err:=service.ServeHTTP(r)
	if err!=nil{
        log.Error("微信第三方开放平台component_verify_ticket获取失败:",err.Error())
        c.Ctx.HTML("error")
        return
	}

    // 将ticket缓存,并在服务重启时取用.
	cache.Set("cachekey_of_wechat3rd_ticket",resp.ComponentVerifyTicket,time.Hour*12)
	err=service.SetTicket(resp.ComponentVerifyTicket)

    if err!=nil{
        log.Error("微信第三方开放平台component_verify_ticket设置失败:",err.Error())
        c.Ctx.HTML("error")
        return
    }
    c.Ctx.HTML("success")

#### 2.4: 使用service获取预授权码与授权链接
    // 获取预授权码并组装链接
    // https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html

    // 方式一: 直接获取授权链接
    authurl,err:=service.AuthUrl(true,"你的授权回调url",wechat3rd.PREAUTH_AUTH_TYPE_MINIAPP,nil)
    
    // 方式二: 你也可以自行获取预授权码并手动拼接授权链接 例如以下代码
	resp:=service.PreAuthCode()
	if !resp.Success(){
        log.Error("获取授权链接失败:",resp.ErrMsg)
        return
	}

    r := url.Values{}

    // 必选参数
	r.Add("component_appid",os.Getenv("WX_OPEN_APP_ID"))
	r.Add("pre_auth_code",resp.PreAuthCode)
	r.Add("redirect_uri","你的回调url")
	r.Add("auth_type",wechat3rd.PREAUTH_AUTH_TYPE_MINIAPP)

    // 网页方式授权：授权注册页面扫码授权
    authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?"
    authUrl += r.Encode()
    println(authUrl)

    // 移动设备方式授权：点击移动端链接快速授权
    r.Add("action","bindcomponent")
    r.Add("no_scan",strconv.Itoa(1))
    authUrl := "https://mp.weixin.qq.com/safe/bindcomponent?"
    authUrl += r.Encode()+"#wechat_redirect"
    println(authUrl)

#### 2.5: 预授权回调处理

    // 预授权链接授权操作之后微信会对回调url发起GET请求.
    // https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
    
    // 获取GET auth_code参数 这里用iris做案例
    authCode := c.Ctx.URLParam("auth_code")

    // 该处service应该是从你的单例方法或者服务池中获取
    rsp:=service.QueryAuth(authCode)
	if !rsp.Success(){
        return errors.New("换取token失败:"+rsp.ErrMsg)
	}
    
    //授权方(也就是小程序持有方,非第三方)令牌信息
    info:=rsp.AuthorizationInfo
    //info.AuthorizerAppid 授权方的appid
	//info.AuthorizerRefreshToken 授权方的刷新token
	//info.AuthorizerAccessToken 授权方的authAccessToken

    //接下来做你想做的,例如缓存授权方的token,存储授权方的刷新token等

	c.Ctx.HTML("授权成功")

### 3.Service方法说明：

    ServeHTTP: 处理消息与事件接收URL的推送 例如:获取component_verify_ticket, 小程序审核, 类目审核等结果推送
    Token: 获取第三方平台的token
    AuthorizerInfo: 获取授权详情
    AuthorizerOption： 获取选项信息
    SetAuthorizerOption： 设置选项
    AuthorizerList： 选项列表
    PostJson： 提交json数据
    PreAuthCode： 获取令牌
    AuthUrl： 获取授权连接
    QueryAuth: 获取授权公众号信息， 注意返回的token,appid等信息需要自行保存，后面代小程序或公众号实现业务时使用
    RefreshToken: 刷新授权用户的token
    AESDecryptData: 用于解密数据

### 4.小程序登陆

    Jscode2session: 获取用户openid , session_key
    AESCBCDecrypt: 用于解密用户数据, 例如解密前端获取手机号时获取的加密信息

### 5.代小程序实现业务

    FastRegisterWeapp: 快速创建小程序
    SearchWeapp: 查询创建任务状态

    * 小程序模板接口 (全部完成)
    GetTemplateDraftList: 获取代码草稿列表
    AddToTemplate: 将草稿添加到代码模板库
    GetTemplateList: 获取代码模板列表
    DeleteTemplate: 删除指定代码模板

**!!!以下接口注意:accessToken(authorizerAccessToken)为授权方token!!!**

    * 基础信息设置 (部分完成)
    GetAccountBasicInfo: 获取基本信息
    ModifyDomain: 设置服务器域名
    SetWebviewDomain: 设置业务域名

    * 小程序用户隐私保护指引(全部完成)
    SetPrivacySetting: 设置小程序用户隐私保护指引
    GetPrivacySetting: 查询小程序用户隐私保护指引
    UploadPrivacyExtFile: 上传小程序用户隐私保护指引

    * 小程序类目管理接口 (全部完成)
    GetMiniProgramAllCategory: 获取可以设置的所有类目
    GetMiniProgramCategory: 获取已设置的所有类目
    GetMiniProgramCategoriesByType: 获取不同主体类型的类目
    AddMiniProgramCategory: 添加类目
    DeleteMiniProgramCategory: 删除类目
    ModifyMiniProgramCategory: 修改类目资质信息
    GetMiniProgramAuditCategory: 获取审核时可填写的类目信息

    * 成员管理 (全部完成)
    BindTester: 绑定体验者
    UnbindTester: 解除绑定体验者
    Memberauth: 获取体验者列表

    * 代码管理 (全部完成)
    Commit: 上传代码
    GetPage: 获取已上传的代码的页面列表
    GetQrcode: 获取体验版二维码
    SubmitAudit: 提交审核
    GetAuditStatus: 查询指定发布审核单的审核状态
    GetLatestAuditStatus: 查询最新一次提交的审核状态
    UndoCodeAudit: 小程序审核撤回
    Release: 发布已通过审核的小程序
    RevertCodeRelease: 小程序版本回退
    GetRevertCodeRelease: 获取可回退的小程序版本
    GrayRelease: 分阶段发布(灰度发布)
    GetGrayReleasePlan: 查询当前分阶段发布详情
    RevertGrayRelease: 取消分阶段发布
    ChangeVisitStatus: 修改小程序服务状态
    GetWeappSupportVersion: 查询当前设置的最低基础库版本及各版本用户占比
    SetWeappSupportVersion: 设置最低基础库版本
    QueryQuota: 查询服务商的当月提审限额（quota）和加急次数
    SpeedupAudit: 加急审核申请

    * 订阅消息设置  (全部完成)
    GetCategory: 获取当前帐号所设置的类目信息
    GetPubTemplateTitles: 获取模板标题列表
    GetPubTemplateKeywords: 获取模板标题下的关键词库
    AddTemplate: 组合模板并添加到个人模板库
    GetTemplate: 获取帐号下的模板列表
    DelTemplate: 删除帐号下的某个模板
    SubscribeSend: 发送订阅消息

    * 支付后获取 Unionid
    GetPaidUnionId: 支付后获取用户 Unionid 接口

    * 素材管理
    MediaUpload: 新增临时素材
    MediaGet: 获取临时素材
    GetMaterial: 获取永久素材

    * 插件管理
    Plugin: 小程序插件管理

## todo
    * 代小程序实现业务(部分完成)
    * 开放平台账号管理
    * 代公众号实现业务