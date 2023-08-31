/*
 *@Author: frank
 *@Date: 2023-08-01 13:40:19
 *@Description: 三方平台配置项
 */

package wxopen

import (
	"os"
	"strings"
)

type Config struct {
	ComponentAppId        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_secret"`
	ComponentVerifyTicket string `json:"component_verify_ticket,optional"`
	Aeskey                string `json:"aeskey"`
	V3key                 string `json:"v3key"`
	RedirectUrl           string `json:"redirect_url"` // 开放平台验证票据回调url
	GatewayHost           string `json:"gateway_host"`
	AppCertPath           string `json:"app_cert_path"`
	PrivateKey            string `json:"private_key,optional"`
	PublicKey             string `json:"public_key,optional"`
	V3PublicKey           string `json:"v3_public_key,optional"`
	WechatpaySerial       string `json:"wechatpay_serial"`
}

var _option Config

func NewOption(o Config) Config {
	certPath := ""
	if o.AppCertPath != "" {
		certPath = o.AppCertPath
	}

	if strings.TrimSpace(certPath) != "" {
		privateKeyBytes, err := os.ReadFile(certPath + "/wechat_private_key.pem")
		if err != nil {
			panic("读取私钥出错，文件不存在！")
		}

		privateKey := string(privateKeyBytes)
		o.PrivateKey = privateKey

		publicKeyBytes, err := os.ReadFile(certPath + "/wechat_public_key.pem")
		if err != nil {
			panic("读取私钥出错，文件不存在！")
		}

		publicKey := string(publicKeyBytes)
		o.PublicKey = publicKey

		v3PublicKeyBytes, err := os.ReadFile(certPath + "/v3_public_key.pem")
		if err != nil {
			panic("读取私钥出错，文件不存在！")
		}
		v3PublicKey := string(v3PublicKeyBytes)
		o.PublicKey = v3PublicKey
	}

	_option = o

	return _option
}

func GetOption() Config {
	return _option
}

func SetOption(config Config) {

}
