package xhttp

import (
	"fbstrans/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

//http方法
func HttpResult(r *http.Request,w http.ResponseWriter,resp interface{},err error)  {

	if err == nil {
		//成功返回
		r:= Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.BAD_REUQEST_ERROR
		errmsg := "服务器繁忙，请稍后再试"

		causeErr := errors.Cause(err) // err类型
		if e,ok := causeErr.(*xerr.CodeError);ok{ //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		}else{
			if gstatus, ok := status.FromError(causeErr);ok{ // grpc err错误
				errmsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ",err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode,errmsg))
	}
}



func HttpDTMResult(r *http.Request,w http.ResponseWriter,resp interface{},err error)  {

	if err == nil {
		//成功返回
		r:= DtmSuccess(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.BAD_REUQEST_ERROR
		errmsg := "服务器繁忙，请稍后再试"

		causeErr := errors.Cause(err) // err类型
		if e,ok := causeErr.(*xerr.CodeError);ok{ //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		}else{
			if gstatus, ok := status.FromError(causeErr);ok{ // grpc err错误
				errmsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ",err)

		httpx.WriteJson(w, http.StatusBadRequest, DtmError(errcode,errmsg))
	}
}




//http 参数错误返回
func ParamErrorResult(r *http.Request,w http.ResponseWriter,err error)  {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQES_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQES_PARAM_ERROR,errMsg))
}
