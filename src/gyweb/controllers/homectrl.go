package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"../models"
	"./common"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

// 导航栏统一入口
func (this *HomeController) Get() {
	op := this.Input().Get("op")
	beego.Info("导航栏统一入口 操作方式", op)
	switch op {
	case "category":
		Tocategory(&(this.Controller), "all", "-1", 0)
	case "gallery":
		ToGallery(this, 1)
	case "about":
		ToAbout(this)
	default:
		tohome(this)
	}
	return

}

//首页
func tohome(this *HomeController) {

	arr := common.GetNavData(0, &(this.Controller))
	picarr := models.GetRollPic()
	hotarr, _, _ := models.GetImagePostsByName("-all", "-BrowseCount", 0, 4)
	this.Data["NavArr"] = arr
	this.Data["RollingPictureArr"] = picarr
	this.Data["HotArr"] = hotarr
	this.TplName = "home.html"
}

/**
发表页面
*/

func (this *HomeController) Publish() {
	arr := common.GetNavData(2, &(this.Controller))
	this.Data["NavArr"] = arr
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "publish.html"
}

/**
图库专用信息集合
*/
type ReGallery struct {
	Title      string
	ImageUrl   string
	ClassLabel string
	TextLabel  string
	Name       string
}

/**
图库
*/
func ToGallery(this *HomeController, page int) {
	arr := common.GetNavData(3, &(this.Controller))
	this.Data["NavArr"] = arr
	p, num, err := models.GetImagePosts(page, 8)

	var data []ReGallery
	if err == nil {
		data = make([]ReGallery, num)
		for i, v := range p {
			s := v.ImageLabel
			s = strings.Replace(s, "web-design", "设计", -1)
			s = strings.Replace(s, "graphic", "风景", -1)
			data[i] = ReGallery{Title: v.Title, ImageUrl: v.Images[0].Url, ClassLabel: v.ImageLabel, TextLabel: s, Name: v.Author.Name}
		}
	}

	this.Data["PostArr"] = data
	this.Data["Currindex"] = page
	this.Data["MaxIndex"] = models.GetImageCountPage()
	this.TplName = "gallery.html"
}

/*
func (this *HomeController) UserCenter() {
	ToUserCenter(&(this.Controller), 0)
}
func (this *HomeController) AdminCenter() {

	name := this.GetSession("uname").(string)
	u := models.GetUserInfo(name)
	p, _, _ := models.GetHotBlogByName("", 4)
	this.Data["User"] = u
	this.Data["PostArr"] = p
	this.TplName = "user/profile.tpl"
}

/**
对外开放的方法，在文章控制器里使用到

func ToUserCenter(this *beego.Controller, page int) {
	arr := common.GetNavData(4, this)
	this.Data["NavArr"] = arr
	uname, ok := this.GetSession("uname").(string)
	if ok {
		p, pageCount, _, err := models.QueryListByUserName(uname, page)
		if err == nil {
			beego.Info("用户中心 页数", pageCount)
			pl := PageLimit{CurrentIndex: page + 1, IndexCount: int(pageCount)}
			pl.IndexArr = make([]int, int(pageCount))
			for i := 0; i < pl.IndexCount; i++ {
				pl.IndexArr[i] = i + 1
			}
			this.Data["PageLimit"] = pl
			this.Data["PostArr"] = p
			this.TplName = "usercenter.html"
			return
		}

	}
	this.Redirect("/", 302)
}*/

/**
关于我们
*/
func ToAbout(this *HomeController) {
	arr := common.GetNavData(5, &(this.Controller))
	this.Data["NavArr"] = arr
	this.TplName = "about.html"
}

//检查用户是否已登录
func CheckUserLogin(this *beego.Controller) {
	/**
	  判断cookie与session
	*/
	uname, _, err := CheckUserCookie(this.Ctx)
	if err {
		v := this.GetSession("uname")
		beego.Info("获取Session", v, "key:", uname)
		if v != nil {
			beego.Info(" 判断cookie与session")
			this.Data["Islogin"] = true
		}
		s := "?type=logout&uname="
		this.Data["LoginOP"] = fmt.Sprint(common.To_login_Page, s, uname)
	} else {
		this.Data["LoginOP"] = fmt.Sprint(common.To_login_Page, "?type=login")
	}
}

func (this *HomeController) Prepare() {
	CheckUserLogin(&(this.Controller))
	//XSRF 之前是全局设置的一个参数,如果设置了那么所有的API请求都会进行验证,但是有些时候API 逻辑是不需要进行验证的,因此现在支持在controller 级别设置屏蔽:
	this.EnableXSRF = false
}
