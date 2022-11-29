package bo

import "github.com/lutasam/gin_admin_sys/biz/vo"

type AddCommodityRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Count int     `json:"count" binding:"required"`
}

type AddCommodityResponse struct{}

type UpdateCommodityRequest struct {
	ID    string  `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"-"`
	Price float32 `json:"price" binding:"-"`
	Count int     `json:"count" binding:"-"`
}

type UpdateCommodityResponse struct{}

type FindCommoditiesRequest struct {
	CurrentPage int    `json:"current_page" binding:"required"`
	PageSize    int    `json:"page_size" binding:"required"`
	SearchName  string `json:"search_name" binding:"-"`
}

type FindCommoditiesResponse struct {
	Total       int               `json:"total"`
	Commodities []*vo.CommodityVO `json:"commodities"`
}

type DeleteCommodityRequest struct {
	ID string `json:"id" binding:"required"`
}

type DeleteCommodityResponse struct{}
