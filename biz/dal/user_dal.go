package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/model"
	"github.com/lutasam/gin_admin_sys/biz/repository"
	"sync"
)

type UserDal struct{}

var (
	userDal     *UserDal
	userDalOnce sync.Once
)

func GetUserDal() *UserDal {
	userDalOnce.Do(func() {
		userDal = &UserDal{}
	})
	return userDal
}

// GetUserByUsername if there is no this username in database, it will return error
func (ins *UserDal) GetUserByUsername(c *gin.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Where("username = ?", username).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Create(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
