欢迎使用 Wechat Easy SDK Go。本SDK非官方实现

# 微信公众号相关
## 自调用
### 获取Access token
所属包：mp  
方法名称：AccessToken  
参数：appId string,secret string
示例：
```
func TestAccessToken(t *testing.T) {
	res, err := AccessToken("wxce4c356a74b76720", "*********")
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("res", res)
}
```