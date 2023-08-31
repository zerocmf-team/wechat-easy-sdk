/**
** @创建时间: 2022/8/15 16:57
** @作者　　: return
** @描述　　:
 */

package base

import (
	"fmt"
	"testing"
)

func TestAccessToken(t *testing.T) {
	res, err := Token("wxce4c356a74b76720", "df5cca115cf3db90736952d51a51a4c7")
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("res", res)
}
