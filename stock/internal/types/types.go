// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type DeductRequest struct {
	GoodsId int64 `json:"goods_id"`
	Num     int64 `json:"num"`
}

type DeductResponse struct {
	GoodsId int64 `json:"goods_id"`
}
