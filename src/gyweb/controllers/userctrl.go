package controllers

import (
	"html/template"

	"../models"
	"./common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"../utils"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	logtype := this.Input().Get("type")
	if logtype == "logout" {
		uname := this.Input().Get("uname")
		if uname != "" {

			this.Ctx.SetCookie("uname", uname, -1, "/")
			this.Ctx.SetCookie("pwd", uname, -1, "/")
			this.DelSession("uname")
			this.DelSession("GategoryManagementType")
			this.DelSession("GategoryManagementOp")
			this.DelSession("SectionType")
			this.DelSession("Section")
			this.DelSession("isadmin")
			this.Redirect("/", 302)
			return
		}

	} else if logtype == "register" {
		this.Data["register"] = true
	} else {
		this.Data["register"] = false
	}
	s := this.Input().Get("from")
	if len(s) > 0 {
		this.SetSession("from", s)
	}
	/**
	跨站请求伪造(Cross-site request forgery)， 简称为 XSRF，是 Web 应用中常见的一个安全问题。前面的链接也详细讲述了 XSRF 攻击的实现方式。
	当前防范 XSRF 的一种通用的方法，是对每一个用户都记录一个无法预知的 cookie 数据，然后要求所有提交的请求（POST/PUT/DELETE）中都必须带有这个
	cookie 数据。如果此数据不匹配 ，那么这个请求就可能是被伪造的。
	*/
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "login.html"
}

func (this *UserController) Post() {
	u := common.UserPost{}
	err := this.ParseForm(&u)
	if err == nil {
		su, isadmin := models.Login(u.Uname, u.Pwd)
		if su {
			Success(this, &u, isadmin)

		}
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["err"] = true
	this.TplName = "login.html"

}

func (this *UserController) Register() {
	u := common.UserPost{}
	err := this.ParseForm(&u)
	if err == nil {
		if models.CheckUser(u.Uname) {
			_, err := models.AddUser(u.Uname, u.Pwd)
			if err == nil {
				Success(this, &u, false)
				return
			}

		}
		this.Data["errmsg"] = "该用户已存在"
	} else {
		this.Data["errmsg"] = "服务器出错"
	}
	this.Data["err"] = true
	this.Data["register"] = true
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "login.html"
	return

}

func Success(this *UserController, u *common.UserPost, isadmin bool) {
	maxTime := 0
	if u.Autologin == "on" {
		maxTime = 1 << 32 - 1
	}
	this.Ctx.SetCookie("uname", u.Uname, maxTime, "/")
	this.Ctx.SetCookie("pwd", u.Pwd, maxTime, "/")

	this.SetSession("uname", u.Uname)
	this.SetSession("isadmin", isadmin)

	from := this.GetSession("from")
	beego.Info("Success 获取来自页面", from)
	if from != nil {
		this.DelSession("from")
		this.Redirect(from.(string), 302)
	} else {
		this.Redirect("/", 302)
	}

	return
}

func CheckUserCookie(cxt *context.Context) (string, string, bool) {

	ck, err := cxt.Request.Cookie("uname")
	if err != nil {
		return "", "", false
	}

	pwd, err1 := cxt.Request.Cookie("pwd")
	if err1 != nil {
		return "", "", false
	}

	return ck.Value, pwd.Value, true
}

func (this *UserController)EditPwdPage() {
	curruname := this.GetSession("uname").(string)
	cu := models.GetUserInfo(curruname)
	this.Data["User"] = cu
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "user/profile-pwd.tpl"

}
func (this *UserController)UpdateUserPwd() {
	/**
	  "uid": $('#uid').val(),
		    "oldpwd": $('#oldpwd').val(),
		    "newpwd": $('#newpwd').val(),
		    "confpwd": $("#confpwd").val()
	 */
	uid := this.Input().Get("uid")
	oldpwd := this.Input().Get("oldpwd")
	newpwd := this.Input().Get("newpwd")
	confpwd := this.Input().Get("confpwd")
	beego.Info("修改密码", uid, oldpwd, newpwd, confpwd)
	re := struct {
		Ok  int
		Msg string
	}{200, "修改成功"}
	if newpwd != confpwd {

		re.Ok = 400
		re.Msg = "两者密码不一样"
		this.Data["json"] = &re
	} else {
		b := models.UpdateUserPwd(uid, oldpwd, newpwd)
		if b {
			this.Data["json"] = &re
		} else {
			re.Ok = 400
			re.Msg = "旧密码错误"
			this.Data["json"] = &re
		}
	}

	this.ServeJSON()
}

func (this *UserController)EditInfoPage() {

	uid := this.Ctx.Input.Param(":id")
	u := models.GetUserInfoByID(uid)
	curruname := this.GetSession("uname").(string)
	cu := models.GetUserInfo(curruname)
	this.Data["Usered"] = u
	this.Data["User"] = cu
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "user/profile-form.tpl"

}

func (this *UserController)EditUserPic() {
	uid := this.Ctx.Input.Param(":id")
	u := models.GetUserInfoByID(uid)
	this.Data["User"] = u
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "user/avatar.tpl"

}

/**

 */
func (this *UserController)UpdateUserPic() {
	uid := this.Ctx.Input.Param(":id")
	u := models.GetUserInfoByID(uid)
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["Title"] = "你并没有选择图片..."
		this.Data["Msg"] = err.Error()
		this.TplName = "error.html"
	} else {
		defer f.Close()
		p := utils.GetUserPicSavePath(h.Filename, u.Name)
		err = this.SaveToFile("file", p)
		if err == nil {
			models.UpdateUserPic(uid, "/" + p)
			this.Redirect(common.To_User_Center, 302)
			return
		} else {
			this.Data["Title"] = "修改头像失败..."
			this.Data["Msg"] = err.Error()
			this.TplName = "error.html"
		}
	}

}
func (this *UserController) UpdateInfo() {
	uid := this.Ctx.Input.Param(":id")
	uform := common.UpdateUserForm{}
	this.ParseForm(&uform)
	beego.Info("更新用户资料", uform)
	err := models.UpdateById(uid, uform)
	if err == nil {
		this.Redirect(common.To_User_Center, 302)
	} else {
		this.Data["Title"] = "修改资料失败..."
		this.Data["Msg"] = err.Error()
		this.TplName = "error.html"
	}
}

func (this *UserController) BrowseInfo() {

	uid := this.Input().Get("uid")
	u := models.GetUserInfoByID(uid)
	p, _, _ := models.QueryListByColumn("Author__Name", u.Name, "-Created", 10, 0)

	this.Data["User"] = u
	this.Data["PostArr"] = p
	this.TplName = "user/user-info.tpl"

}
/**
用户中心
*/
func (this *UserController) UserCenter() {
	ToUserCenter(&(this.Controller), 0)
}

func ToUserCenter(this *beego.Controller, page int) {

	name := this.GetSession("uname").(string)
	u := models.GetUserInfo(name)
	p, _, _ := models.GetHotBlogByName(name, 5)
	num, bc := models.GetGategoryCountAndBrowseCount(name)

	this.Data["CateCount"] = num
	this.Data["BrowseCount"] = bc

	this.Data["User"] = u
	this.Data["PostArr"] = p
	this.TplName = "user/profile.tpl"
}

func (this *UserController)UserManage() {
	name := this.GetSession("uname").(string)
	u := models.GetUserInfo(name)
	this.Data["User"] = u
	uarr, allnum, _ := models.GetUserList(0)

	px := allnum / 15
	x := allnum % 15
	if x > 0 {
		px += 1
	}

	this.Data["Count"] = allnum
	//分页处理
	pl := PageLimit{CurrentIndex: 0 + 1, IndexCount: int(px)}
	pl.IndexArr = make([]int, int(px))
	for i := 0; i < pl.IndexCount; i++ {
		pl.IndexArr[i] = i + 1
	}
	this.Data["PageLimit"] = pl
	this.Data["UserArr"] = uarr

	this.TplName = "user/profile-usermanage.tpl"
}