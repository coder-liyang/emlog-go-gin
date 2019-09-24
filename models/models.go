package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var Orm *gorm.DB
func ClientDb() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/liyangweb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("链接数据库错误:", err.Error())
		panic("数据库链接失败")
	}
	//表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "e_" + defaultTableName
	}
	//关闭表名的复数形式
	db.SingularTable(true)

	Orm = db
	fmt.Println("数据库链接成功", db)
}
