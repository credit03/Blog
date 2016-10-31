package main

import (
	_ "./routers"
	"./models"
	"./initcreate"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {

	models.RegisterDB()
	initcreate.InitTemplateMaps()

}
func main() {
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	orm.RunSyncdb("db2", false, true)
	models.AutoCreateAdmin()
	models.AutoCreateRollPic()
	beego.Run()
}

