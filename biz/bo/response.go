package bo

import "github.com/lutasam/gin_admin_sys/biz/vo"

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PingResponse struct {
	Pong string `json:"pong"`
}

type HelloResponse struct {
	Hello string `json:"hello"`
}

/*
=================================================
  - Login Module

==================================================
*/

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterResponse struct{}

type ActiveResponse struct{}

/*
=================================================
  - Commodity Module

==================================================
*/
type AddCommodityResponse struct{}

type UpdateCommodityResponse struct{}

type FindAllCommoditiesResponse struct {
	Total        int               `json:"total"`
	CommodityVOs []*vo.CommodityVO `json:"commodity_vos"`
}

type DeleteCommodityResponse struct{}
