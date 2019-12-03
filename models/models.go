package models

import (
	"emlog-go-gin/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Orm *gorm.DB

func ClientDb() {
	username := config.GetConfigs().Mysql.Username
	password := config.GetConfigs().Mysql.Password
	host := config.GetConfigs().Mysql.Host
	dbname := config.GetConfigs().Mysql.Dbname
	charset := config.GetConfigs().Mysql.Charset
	dbLinkTpl := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbname,
		charset,
	)
	db, err := gorm.Open("mysql", dbLinkTpl)
	if err != nil {
		fmt.Println("链接数据库错误:", err.Error())
		panic("数据库链接失败")
	}
	//表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return "e_" + defaultTableName
		return config.GetConfigs().Mysql.Prefix + defaultTableName
	}
	//关闭表名的复数形式
	db.SingularTable(true)
	Orm = db
	if config.GetConfigs().Mysql.Debug {
		Orm = Orm.LogMode(true).Debug()
	}
	fmt.Println("数据库链接成功", db)
}
