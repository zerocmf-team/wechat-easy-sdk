/**
** @创建时间: 2022/8/15 17:13
** @作者　　: return
** @描述　　:
 */

package data

type Response struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
