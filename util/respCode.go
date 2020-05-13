package util

// 请求状态码
const (
	RECODE_OK      = 1
	RECODE_FAIL    = 0
	RECODE_UNLOGIN = -1
)

// 业务状态码
const (
	RESPMSG_OK = "1"
	RESPMSG_FAIL = "0"
	ERROR_UNLOGIN = "ERROR_UNLOGIN"

	RECODE_UNKNOWN = ""
	// 管理员
)

var recodeText = map[string]string{
	RESPMSG_OK: "成功",
	RESPMSG_FAIL: "失败",
	ERROR_UNLOGIN: "未登录，请先登录",
}

func Recode2Text(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWN]
}