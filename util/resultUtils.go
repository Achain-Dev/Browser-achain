package util

func GetWebResultSuccess() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 200
	result["msg"] = "OK"
	return result
}

func GetWebResultFail() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 1001
	result["msg"] = "系统错误，请稍后再试"
	return result
}
