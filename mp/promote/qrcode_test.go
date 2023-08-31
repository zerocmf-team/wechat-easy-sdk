/**
** @创建时间: 2022/8/25 08:12
** @作者　　: return
** @描述　　:
 */

package promote

import (
	"fmt"
	"testing"
)

func TestQrcode_Create(t *testing.T) {
	qrcode := new(Qrcode)
	res, err := qrcode.Create("aQjcG4Dbr2wbLKauJDA8AeSRpPB07hXvfYqlNQ4BHnvYFqS4ZB2OxL_VT4QgS6n7Eewd9cfU6ortczxmd0b3iAQceGiA-w3cDTlqnWGCZFAh5PiajXBVJlb5c7Q29DGbhocfXxiF4H34wCrBXRGeAGAGAP", 120,qrcode.QrStrScene(),qrcode.WithSceneStr("111111111"))
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("res", res)
}
