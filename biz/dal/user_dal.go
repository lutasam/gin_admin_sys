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

// TakeUserByEmail if there is no this email in database, it will return error
func (ins *UserDal) TakeUserByEmail(c *gin.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Create(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) UpdateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Model(user).Save(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
