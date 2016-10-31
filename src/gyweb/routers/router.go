package routers

import (
	"fmt"

	"../controllers"
	"../controllers/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	//导航栏
	beego.Router("/", &controllers.HomeController{})
	beego.Router(common.To_Publish, &controllers.HomeController{}, "*:Publish")

	//beego.Router(common.To_Admin_Center, &controllers.HomeController{}, "*:AdminCenter")
	//用户操作
	beego.Router(common.To_login_Page, &controllers.UserController{}, "*:Get")
	beego.Router(common.To_login, &controllers.UserController{}, "*:Post")
	beego.Router(common.To_register, &controllers.UserController{}, "*:Register")
	beego.Router(common.To_userinfo_browse, &controllers.UserController{}, "*:BrowseInfo")
	beego.Router(common.To_userinfo_edit, &controllers.UserController{}, "*:EditInfoPage")
	beego.Router(common.To_userinfo_update, &controllers.UserController{}, "*:UpdateInfo")

	beego.Router(common.To_userpic_edit, &controllers.UserController{}, "*:EditUserPic")
	beego.Router(common.To_userpic_update, &controllers.UserController{}, "*:UpdateUserPic")

	beego.Router(common.To_userpwd_edit, &controllers.UserController{}, "*:EditPwdPage")
	beego.Router(common.To_userpwd_update, &controllers.UserController{}, "*:UpdateUserPwd")
	//上传
	beego.Router(common.Image_Upload, &controllers.UploadController{}, "*:ImageUpload")
	//备用视频接口，
	beego.Router(common.Viedo_Upload, &controllers.UploadController{}, "*:VideoUpload")

	//文章操作
	beego.Router(common.To_Category_Add, &controllers.CategoryController{}, "*:Add")
	beego.Router(common.To_Category_Delete, &controllers.CategoryController{}, "*:Delete")
	beego.Router(common.To_Category_Edit, &controllers.CategoryController{}, "*:ToEditPage")
	beego.Router(common.To_Category_Updtae, &controllers.CategoryController{}, "*:EditUpdate")
	beego.Router(common.To_Category_Browse, &controllers.CategoryController{}, "*:BrowseCategory")
	beego.Router(common.AutoTip, &controllers.CategoryController{}, "*:AjaxAutoTip")
	//自动url
	beego.AutoRouter(&controllers.CategoryController{})

	//后台链接
	beego.Router(common.To_User_Center, &controllers.UserController{}, "*:UserCenter")
	beego.Router(common.To_Center_Gategory_manage, &controllers.CategoryController{}, "*:GateGoryCenter")
	beego.Router(common.To_RollingPic_Set, &controllers.CategoryController{}, "*:SetPollingPic")
	beego.Router(common.To_User_Manage, &controllers.UserController{}, "*:UserManage")

	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	Filter()
}

func Filter() {

	var FilterUser = func(ctx *context.Context) {

		_, ok := ctx.Input.Session("uname").(string)
		beego.Info("授权操作:", ok, ctx.Request.RequestURI)
		if !ok && ctx.Request.RequestURI != "/myajax/*" {
			url := fmt.Sprint(common.To_login_Page, "?from=", ctx.Request.RequestURI)
			ctx.Redirect(302, url)
		}
	}

	//授权操作
	beego.InsertFilter("/myajax/*", beego.BeforeRouter, FilterUser)

}
