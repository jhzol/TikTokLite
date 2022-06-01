//package repository
package common

import (
	"TikTokLite/log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DataBase *gorm.DB

func InitDatabase() {
	var err error
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	log.Info(args)
	DataBase, err = gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	log.Infof("connect database success,user:%s,database:%s", username, database)
}

func GetDB() *gorm.DB {
	return DataBase
}
func CloseDataBase() {
	DataBase.Close()
}
