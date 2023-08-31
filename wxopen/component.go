/**
** @创建时间: 2021/4/20 5:37 下午
** @作者　　: return
** @描述　　:
 */

package wxopen

import (
	"encoding/json"
	"errors"
	"github.com/zerocmf/wechatEasySdk/data"
	"github.com/zerocmf/wechatEasySdk/util"
)

type Component struct{}

type AccessToken struct {
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	data.Response
}

type funcscopeCategory struct {
	Id int `json:"id"`
}

type funcInfo struct {
	FuncscopeCategory funcscopeCategory `json:"funcscope_category"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/4/20 22:13:19
 * @Param
 * @return
 **/

type ComponentAccessToken struct {
	ComponentAppId        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

func (rest *Component) ComponentAccessToken(bizContent ComponentAccessToken) (result AccessToken, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	data, err := util.PostJson(url, bizContent)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

type PreAuthCode struct {
	ComponentAppid       string `json:"component_appid"`
	ComponentAccessToken string `json:"-"`
}

type PreAuthCodeResult struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 令牌（component_access_token）是第三方平台接口的调用凭据。令牌的获取是有限制的，每个令牌的有效期为 2 小时，请自行做好令牌的管理，在令牌快过期时（比如1小时50分），重新调用接口获取。
 * @Date 2021/4/20 17:41:26
 * @Param
 * @return
 **/

func (rest *Component) PreAuthCode(bizContent PreAuthCode) (result PreAuthCodeResult, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + bizContent.ComponentAccessToken
	var bytes []byte
	bytes, err = util.PostJson(url, bizContent)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return PreAuthCodeResult{}, err
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询授权
 * @Date 2021/4/21 1:10:7
 * @Param
 * @return
 **/

type QueryAuth struct {
	ComponentAccessToken string `json:"-"`
	ComponentAppId       string `json:"component_appid"`
	AuthorizationCode    string `json:"authorization_code"`
}

type authorizationInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfo               []funcInfo `json:"func_info"`
}

type AuthorizationResult struct {
	AuthorizationInfo authorizationInfo `json:"authorization_info"`
	data.Response
}

func (rest *Component) QueryAuth(bizContent QueryAuth) (result AuthorizationResult, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + bizContent.ComponentAccessToken
	var bytes []byte
	bytes, err = util.PostJson(url, bizContent)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 在公众号/小程序接口调用令牌（authorizer_access_token）失效时，可以使用刷新令牌（authorizer_refresh_token）获取新的接口调用令牌。
 * @Date 2021/4/22 22:36:29
 * @Param
 * @return
 **/

type AuthorizerToken struct {
	ComponentAccessToken   string `json:"-"`
	ComponentAppID         string `json:"component_appid"`
	AuthorizerAppID        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type AuthorizerTokenResult struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	data.Response
}

func (rest *Component) AuthorizerToken(bizContent AuthorizerToken) (result AuthorizerTokenResult, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=" + bizContent.ComponentAccessToken
	var bytes []byte
	bytes, err = util.PostJson(url, bizContent)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 第三方平台开发者的服务器使用登录凭证（code）以及第三方平台的 component_access_token 可以代替小程序实现登录功能 获取 session_key 和 openid。其中 session_key 是对用户数据进行加密签名的密钥。为了自身应用安全，session_key 不应该在网络上传输。
 * @Date 2021/4/21 10:17:21
 * @Param
 * @return
 **/

type Code2Session struct {
	AppId                string `json:"app_id"`
	Secret               string `json:"secret"`
	JsCode               string `json:"js_code"`
	ComponentAppid       string `json:"component_appid"`
	ComponentAccessToken string `json:"component_access_token"`
}

type Code2SessionResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid,omitempty"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 第三方平台开发者的服务器使用登录凭证（code）以及第三方平台的 component_access_token 可以代替小程序实现登录功能 获取 session_key 和 openid。其中 session_key 是对用户数据进行加密签名的密钥。为了自身应用安全，session_key 不应该在网络上传输。
 * @Date 2021/4/21 10:25:28
 * @Param
 * @return
 **/

func (rest *Component) Code2session(bizContent Code2Session) (result Code2SessionResult, err error) {
	url := "https://api.weixin.qq.com/sns/component/jscode2session?appid=" + bizContent.AppId + "&js_code=" + bizContent.JsCode + "&grant_type=authorization_code&component_appid=" + bizContent.ComponentAppid + "&component_access_token=" + bizContent.ComponentAccessToken
	var bytes []byte
	bytes, err = util.PostJson(url, nil)
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}
	return
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 快速创建小程序
 * @Date 2021/5/6 15:48:43
 * @Param
 * @return
 **/

type FastRegisterWeapp struct {
	ComponentAccessToken string `json:"-"`
	Name                 string `json:"name"`
	Code                 string `json:"code"`
	CodeType             int    `json:"code_type"`
	LegalPersonaWechat   string `json:"legal_persona_wechat"`
	LegalPersonaName     string `json:"legal_persona_name"`
	ComponentPhone       string `json:"component_phone"`
}

func (rest *Component) FastRegisterWeapp(bizContent FastRegisterWeapp) (result data.Response, err error) {
	if bizContent.ComponentAccessToken == "" {
		err = errors.New("ComponentAccessToken不存在")
		return
	}
	url := "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=create&component_access_token=" + bizContent.ComponentAccessToken
	var data []byte
	data, err = util.PostJson(url, bizContent)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询申请单
 * @Date 2021/5/12 15:54:29
 * @Param
 * @return
 **/

func (rest *Component) FastRegisterWeappSearch(ComponentAccessToken string, bizContent map[string]interface{}) (result data.Response) {
	url := "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=search&component_access_token=" + ComponentAccessToken
	data, _ := util.PostJson(url, bizContent)
	json.Unmarshal(data, &result)
	return result
}
