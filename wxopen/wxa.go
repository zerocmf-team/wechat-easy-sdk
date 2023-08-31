package wxopen

import (
	"encoding/json"
	"github.com/zerocmf/wechatEasySdk/util"

	"github.com/zerocmf/wechatEasySdk/data"
)

type Wxa struct {
}

type ModifyDomain struct {
	AuthorizerAccessToken string   `json:"-"`
	Action                string   `json:"action"`
	RequestDomain         []string `json:"requestdomain,omitempty"`
	WSRequestDomain       []string `json:"wsrequestdomain,omitempty"`
	UploadDomain          []string `json:"uploaddomain,omitempty"`
	DownloadDomain        []string `json:"downloaddomain,omitempty"`
	UDPDomain             []string `json:"udpdomain,omitempty"`
	TCPDomain             []string `json:"tcpdomain,omitempty"`
}

type ModifyDomainResult struct {
	data.Response
	RequestDomain          []string `json:"requestdomain"`
	WsRequestDomain        []string `json:"wsrequestdomain"`
	UploadDomain           []string `json:"uploaddomain"`
	DownloadDomain         []string `json:"downloaddomain"`
	InvalidRequestDomain   []string `json:"invalid_requestdomain"`
	InvalidWsRequestDomain []string `json:"invalid_wsrequestdomain"`
	InvalidUploadDomain    []string `json:"invalid_uploaddomain"`
	InvalidDownloadDomain  []string `json:"invalid_downloaddomain"`
}

func (rest *Wxa) ModifyDomain(bizContent ModifyDomain) (result ModifyDomainResult, err error) {
	url := "https://api.weixin.qq.com/wxa/modify_domain?access_token=" + bizContent.AuthorizerAccessToken
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
 * @Description 设置业务域名
 * @Date 2021/5/5 9:4:2
 * @Param
 * @return
 **/

type SetWebViewDomain struct {
	AuthorizerAccessToken string   `json:"-"`
	Action                string   `json:"action"`
	WebViewDomain         []string `json:"webviewdomain,omitempty"`
}

type WebViewDomainResult struct {
	data.Response
	WebviewDomain []string `json:"webviewdomain"`
}

func (rest *Wxa) SetWebViewDomain(bizContent SetWebViewDomain) (result WebViewDomainResult, err error) {
	url := "https://api.weixin.qq.com/wxa/setwebviewdomain?access_token=" + bizContent.AuthorizerAccessToken
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
