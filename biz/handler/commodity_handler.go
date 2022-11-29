package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/service"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

type CommodityController struct{}

func RegisterCommodityRouter(r *gin.RouterGroup) {
	commodityController := &CommodityController{}
	{
		r.POST("/add_commodity", commodityController.AddCommodity)
		r.POST("/update_commodity", commodityController.UpdateCommodity)
		r.POST("/delete_commodity", commodityController.DeleteCommodity)
		r.GET("/find_commodities", commodityController.FindCommodities)
	}
}

func (ins *CommodityController) AddCommodity(c *gin.Context) {
	req := &bo.AddCommodityRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetCommodityService().AddCommodity(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *CommodityController) UpdateCommodity(c *gin.Context) {
	req := &bo.UpdateCommodityRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetCommodityService().UpdateCommodity(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *CommodityController) DeleteCommodity(c *gin.Context) {
	req := &bo.DeleteCommodityRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetCommodityService().DeleteCommodity(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *CommodityController) FindCommodities(c *gin.Context) {
	currentPage, err := utils.StringToInt(c.Query("current_page"))
	pageSize, err := utils.StringToInt(c.Query("page_size"))
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	req := &bo.FindCommoditiesRequest{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		SearchName:  c.Query("search_name"),
	}
	resp, err := service.GetCommodityService().FindCommodities(c, req)
	if err != nil {
		utils.ResponseServerError(c, err.(common.Error))
		return
	}
	utils.ResponseSuccess(c, resp)
}
