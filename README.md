微信公众平台-第三方平台（简称第三方平台）开放给所有通过开发者资质认证后的开发者使用。在得到公众号或小程序运营者（简称运营者）授权后，第三方平台开发者可以通过调用微信开放平台的接口能力，为公众号或小程序的运营者提供账号申请、小程序创建、技术开发、行业方案、活动营销、插件能力等全方位服务。同一个账号的运营者可以选择多家适合自己的第三方为其提供产品能力或委托运营。

### 项目基于 [owen-gxz/open-wechat](https://github.com/owen-gxz/open-wechat) 做了减法,并补充使用说明.

### 主要完成了微信开放平台第三方平台的[接口说明部分](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_verify_ticket.html)

### 1: 引入
    go get -u github.com/l306287405/wechat3rd@master

### 2: 使用NewService方法来创建一个service
    * Config: 配置信息
    * TicketServer: 保存微信传输的ticket信息接口
    * AccessTokenServer: 获取第三方平台的token接口
    * WechatErrorer: 错误信息的处理

#### Service方法说明：
    * AddHander: 
        用于微信时间推送的处理方法(unauthorized,updateauthorized,authorized,component_verify_ticket)
        方法会接收context
    * ServeHTTP: 处理推送事件的
    * Token: 获取第三方平台的token
    * AuthorizerInfo: 获取授权详情
    * AuthorizerOption： 获取选项信息
    * SetAuthorizerOption： 设置选项
    * AuthorizerList： 选项列表
    * PostJson： 提交json数据
    * PreAuthCode： 获取令牌
    * AuthUrl： 获取授权连接
    * QueryAuth: 获取授权公众号信息， 注意返回的token,appid等信息需要自行保存，后面带公众号实现业务时使用
    * RefreshToken: 刷新授权用户的token

## todo 
    * 开放平台账号管理
    * 代公众号实现业务
    * 代小程序实现业务

### 微信公众号接口不会涉及


