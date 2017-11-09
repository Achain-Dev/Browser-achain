package common

type ResponseType struct {

    data interface{}
	code int   // return code
	msg string // return msg
}

// get if response fail
func GetWebResultFail() (ResponseType)  {
	return ResponseType{code:1001,msg:"system error,please try later"}
}

// get if response success
func GetWebResultSuccess() (ResponseType)  {
	return ResponseType{code:200,msg:"success"}
}
