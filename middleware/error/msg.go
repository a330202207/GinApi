package error

var MsgFlags = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "系统错误",
	UNPASS:         "未通过验证",
	NEED_LOGIN:     "需要登录",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_EXIST_USER:     "用户已经存在了",
	ERROR_NOT_EXIST_USER: "用户名或密码错误",
	ERROR_DISABLE_USER:   "用户已被禁止，请联系管理员",
}

//获取信息码
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
