/**
** @创建时间: 2022/8/15 16:47
** @作者　　: return
** @描述　　:
 */

package base

import (
	"encoding/json"

	"github.com/zerocmf/wechatEasySdk/data"
	"github.com/zerocmf/wechatEasySdk/util"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	data.MpResponse
}

func Token(appid string, secret string) (response TokenResponse, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
	res, err := util.Request("GET", url, nil, nil)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}
