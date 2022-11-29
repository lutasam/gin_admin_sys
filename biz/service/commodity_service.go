package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/dal"
	"github.com/lutasam/gin_admin_sys/biz/model"
	"github.com/lutasam/gin_admin_sys/biz/utils"
	"github.com/lutasam/gin_admin_sys/biz/vo"
	"sync"
)

type CommodityService struct{}

var (
	commodityService     *CommodityService
	commodityServiceOnce sync.Once
)

func GetCommodityService() *CommodityService {
	commodityServiceOnce.Do(func() {
		commodityService = &CommodityService{}
	})
	return commodityService
}

func (ins *CommodityService) AddCommodity(c *gin.Context, req *bo.AddCommodityRequest) (*bo.AddCommodityResponse, error) {
	if req.Name == "" || req.Count < 0 || req.Price < 0 {
		return nil, common.USERINPUTERROR
	}
	err := dal.GetCommodityDal().AddCommodity(c, &model.Commodity{
		ID:    utils.GenerateCommodityID(),
		Name:  req.Name,
		Price: req.Price,
		Count: req.Count,
	})
	if err != nil {
		return nil, err
	}
	return &bo.AddCommodityResponse{}, nil
}

func (ins *CommodityService) UpdateCommodity(c *gin.Context, req *bo.UpdateCommodityRequest) (*bo.UpdateCommodityResponse, error) {
	if req.ID == "" || req.Name == "" || req.Price < 0 || req.Count < 0 {
		return nil, common.USERINPUTERROR
	}
	id, err := utils.StringToUint64(req.ID)
	if err != nil {
		return nil, common.USERINPUTERROR
	}
	commodity, err := dal.GetCommodityDal().TakeCommodityByID(c, id)
	if err != nil {
		return nil, err
	}
	commodity.Name = req.Name
	commodity.Price = req.Price
	commodity.Count = req.Count
	err = dal.GetCommodityDal().UpdateCommodity(c, commodity)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateCommodityResponse{}, nil
}

func (ins *CommodityService) DeleteCommodity(c *gin.Context, req *bo.DeleteCommodityRequest) (*bo.DeleteCommodityResponse, error) {
	if req.ID == "" {
		return nil, common.USERINPUTERROR
	}
	id, err := utils.StringToUint64(req.ID)
	if err != nil {
		return nil, common.USERINPUTERROR
	}
	exist, err := isCommodityExist(c, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, common.USERINPUTERROR
	}
	err = dal.GetCommodityDal().DeleteCommodity(c, id)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteCommodityResponse{}, nil
}

func (ins *CommodityService) FindCommodities(c *gin.Context, req *bo.FindCommoditiesRequest) (*bo.FindCommoditiesResponse, error) {
	if req.CurrentPage <= 0 || req.PageSize <= 0 {
		return nil, common.USERINPUTERROR
	}
	commodities, err := dal.GetCommodityDal().FindCommodity(c, req.CurrentPage, req.PageSize, req.SearchName)
	if err != nil {
		return nil, err
	}
	commodityVOs := convertToCommodityVOs(commodities)
	return &bo.FindCommoditiesResponse{
		Total:        len(commodityVOs),
		CommodityVOs: commodityVOs,
	}, nil
}

func convertToCommodityVOs(commodities []*model.Commodity) []*vo.CommodityVO {
	var commodityVOs []*vo.CommodityVO
	for _, commodity := range commodities {
		commodityVOs = append(commodityVOs, &vo.CommodityVO{
			ID:    utils.Uint64ToString(commodity.ID),
			Name:  commodity.Name,
			Price: commodity.Price,
			Count: commodity.Count,
		})
	}
	return commodityVOs
}
