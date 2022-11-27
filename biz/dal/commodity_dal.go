package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/model"
	"github.com/lutasam/gin_admin_sys/biz/repository"
	"sync"
)

type CommodityDal struct{}

var (
	commodityDal     *CommodityDal
	commodityDalOnce sync.Once
)

func GetCommodityDal() *CommodityDal {
	commodityDalOnce.Do(func() {
		commodityDal = &CommodityDal{}
	})
	return commodityDal
}

func (ins *CommodityDal) AddCommodity(c *gin.Context, commodity *model.Commodity) error {
	err := repository.GetDB().WithContext(c).Table(commodity.TableName()).Create(commodity).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *CommodityDal) UpdateCommodity(c *gin.Context, commodity *model.Commodity) error {
	err := repository.GetDB().WithContext(c).Model(commodity).Save(commodity).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *CommodityDal) DeleteCommodity(c *gin.Context, id uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Commodity{}.TableName()).
		Where("id = ?", id).Delete(&model.Commodity{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *CommodityDal) FindAllCommodity(c *gin.Context) ([]*model.Commodity, error) {
	var commodities []*model.Commodity
	err := repository.GetDB().WithContext(c).Table(model.Commodity{}.TableName()).Find(&commodities).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return commodities, nil
}

// TakeCommodityByID if id does not exist, it will return error
func (ins *CommodityDal) FindCommodityByID(c *gin.Context, id uint64) (*model.Commodity, error) {
	commodity := &model.Commodity{}
	err := repository.GetDB().WithContext(c).Table(commodity.TableName()).Where("id = ?", id).Find(commodity).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return commodity, nil
}
