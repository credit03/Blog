package initcreate

import (
	"github.com/astaxie/beego"
	"../utils"
	"../models"
)

func InitTemplateMaps() {
	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	beego.AddFuncMap("getCurrentTime", utils.GetCurrentTime)
	beego.AddFuncMap("getTopNo", utils.GetNo)

	beego.AddFuncMap("getOneImage", models.GetOneImageById)
	beego.AddFuncMap("getImages", models.GetImages)
	beego.AddFuncMap("getCurrentUserIsAdmin", models.CheckIsAdmin)
}
