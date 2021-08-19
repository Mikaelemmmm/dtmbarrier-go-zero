package xhttp

type ResponseSuccessBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode int, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}


type DTMResponseSuccessBean struct {
	DtmResult string `json:"dtm_result"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DtmSuccess(data interface{}) *DTMResponseSuccessBean {
	return &DTMResponseSuccessBean{"SUCCESS",200,"OK",data}
}

type DTMResponseErrorBean struct {
	DtmResult string `json:"dtm_result"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}
func DtmError(errCode int, errMsg string) *DTMResponseErrorBean {
	 return &DTMResponseErrorBean{"FAILURE",errCode, errMsg}
}


