package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/model"
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

func (ins *UserDal) GetUserByUsername(c *gin.Context, username string) (*model.User, error) {
	user := &model.User{}
	//err := repository.GetDB().WithContext(c).Where("username = ?", username).Find(user).Error
	//if err != nil {
	//	return nil, common.DATABASEERROR
	//}
	//if user.ID == 0 {
	//	return nil, common.USERDOESNOTEXIST
	//}
	user.ID = 123456
	user.Username = "admin"
	user.Password = "admin"
	return user, nil
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	//err := repository.GetDB().WithContext(c).Create(user).Error
	//if err != nil {
	//	return common.DATABASEERROR
	//}
	return nil
}
