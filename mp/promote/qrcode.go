/**
** @创建时间: 2022/8/15 17:43
** @作者　　: return
** @描述　　:
 */

package promote

import (
	"bytes"
	"encoding/json"

	"github.com/zerocmf/wechatEasySdk/data"
	"github.com/zerocmf/wechatEasySdk/util"
)

type scene struct {
	SceneId  string `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type Qrcode struct {
	ExpireSeconds int    `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene scene `json:"scene"`
	} `json:"action_info"`
}

type QrOption struct {
	f func(*Optional)
}

type Optional struct {
	SceneId  string `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type ActionName struct {
	f func(ao *ActionOptional)
}

type ActionOptional struct {
	ActionName string `json:"action_name"`
}

func (qrcode *Qrcode) QrScene() ActionName {
	return ActionName{func(o *ActionOptional) {
		o.ActionName = "QR_SCENE"
	}}
}

func (qrcode *Qrcode) QrStrScene() ActionName {
	return ActionName{func(o *ActionOptional) {
		o.ActionName = "QR_STR_SCENE"
	}}
}

func (qrcode *Qrcode) QrLimitScene() ActionName {
	return ActionName{func(o *ActionOptional) {
		o.ActionName = "QR_LIMIT_SCENE"
	}}
}

func (qrcode *Qrcode) WithSceneId(SceneId string) QrOption {
	return QrOption{func(o *Optional) {
		o.SceneId = SceneId
	}}
}

func (qrcode *Qrcode) WithSceneStr(SceneStr string) QrOption {
	return QrOption{func(o *Optional) {
		o.SceneStr = SceneStr
	}}
}

type CreateResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	data.MpResponse
}

func (qrcode *Qrcode) Create(token string, ExpireSeconds int, an ActionName, ops QrOption) (response CreateResponse, err error) {
	opt := new(Optional)
	ops.f(opt)

	aot := new(ActionOptional)
	an.f(aot)

	params := Qrcode{
		ExpireSeconds: ExpireSeconds,
		ActionName:    aot.ActionName,
	}

	if opt.SceneId != "" {
		params.ActionInfo.Scene.SceneId = opt.SceneId
	}
	if opt.SceneStr != "" {
		params.ActionInfo.Scene.SceneStr = opt.SceneStr
	}

	bytesBody, _ := json.Marshal(params)

	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token
	res, err := util.Request("POST", url, bytes.NewBuffer([]byte(bytesBody)), nil)
	if err != nil {
		return
	}

	json.Unmarshal(res, &response)

	return
}
