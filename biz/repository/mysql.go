package repository

import (
	"fmt"
	"github.com/lutasam/gin_admin_sys/biz/model"
	"github.com/lutasam/gin_admin_sys/biz/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		utils.GetConfigResolve().GetConfigString("mysql.user"),
		utils.GetConfigResolve().GetConfigString("mysql.password"),
		utils.GetConfigResolve().GetConfigString("mysql.address"),
		utils.GetConfigResolve().GetConfigString("mysql.port"),
		utils.GetConfigResolve().GetConfigString("mysql.dbname"),
		utils.GetConfigResolve().GetConfigString("mysql.config"))), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(model.User{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
