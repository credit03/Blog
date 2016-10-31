package controllers

import (
	"strconv"

	"../models"
	"./common"
	"../utils"
	"github.com/astaxie/beego"
	"time"
)

/**
管理文章的增加，修改，删除,单一文章浏览操作
*/

type CategoryController struct {
	beego.Controller
}

type PageLimit struct {
	CurrentIndex int
	IndexCount   int
	IndexArr     []int
}

/**
添加文章
*/
func (this *CategoryController) Add() {

	d := &common.PublishForm{}
	err := this.ParseForm(d)
	if err == nil {
		beego.Info("表单内容videourl,", d.Usevideo)
		beego.Info("添加文章,", time.Now().Format("2006-01-02 15:04:05"))
		_, err = models.CategoryAdd(d, this.GetSession("uname").(string))
	}
	this.Redirect(common.To_Category, 302)
	return
}

/**
更新文章
*/
func (this *CategoryController) EditUpdate() {
	d := &common.PublishForm{}
	err := this.ParseForm(d)
	if err == nil {
		name := this.GetSession("uname").(string)
		_, err = models.UdpateCateGory(d, name)
		if err == nil {
			this.Redirect(common.To_Center_Gategory_manage, 302)
			return
		} else {
			this.Data["Title"] = "修改文章失败..."
			this.Data["Msg"] = "你没有权限修改"
			this.TplName = "error.html"
		}
	}

}

/**
删除文章
*/
func (this *CategoryController) Delete() {
	id := this.Input().Get("id")
	name := this.GetSession("uname").(string)
	//防止A帐号删除B帐号的文章
	err := models.DeletePostByID(id, name)
	if err == nil {
		ToUserCenter(&(this.Controller), 0)
	} else {
		this.Data["Title"] = "删除文章失败..."
		this.Data["Msg"] = "文章Id=" + id
		this.TplName = "error.html"
	}
}

/**
跳转到修改文章页面
*/
func (this *CategoryController) ToEditPage() {

	//导航栏
	arr := common.GetNavData(-2, &(this.Controller))
	this.Data["NavArr"] = arr

	id := this.Input().Get("id")
	name := this.GetSession("uname").(string)
	p, err := models.GetOneByUserName("Id", id, name, false)
	if err == nil {
		s := strconv.Itoa(int(p.Id))
		data := &common.PublishForm{Id: s, Title: p.Title, Content: p.Content, ContentType: strconv.Itoa((int(p.ContentType))), Sectiontype: strconv.Itoa((p.SectionType)), Describe: p.Describe}
		if p.ContentType == 1 || p.ContentType == 2 {
			data.ImageLabel = p.ImageLabel
			str := ""
			for _, v := range p.Images {
				str += v.Url + "#substr#"
			}
			data.Useimage = str
		}
		this.Data["Post"] = data
		this.TplName = "edit.html"

	} else {
		this.Data["Title"] = "权限问题或已下线"
		this.Data["Msg"] = "对不起，你没有权限修改该文章或管理员已下线该文章"
		this.TplName = "error.html"
	}

}

func (this *CategoryController) GateGoryCenter() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("获取ID", id)

	o := this.Input().Get("op")
	t := this.Input().Get("type")
	beego.Info("GateGoryCenter...操作类型", t, "操作方式", o)
	GategoryManagement(this, t, o, 0)

}

/**
后台管理中心，登录操作
 */
func GategoryManagement(this *CategoryController, types string, op string, page int) {

	name := this.GetSession("uname").(string)
	u := models.GetUserInfo(name)
	this.Data["User"] = u

	//分页下标
	var index int64

	switch types {

	case "ContentType":
		//获取类型博客
		p, allnum, _, err := models.QueryByContentTypeAjax(&u, op, page)
		if err == nil {
			this.Data["PostArr"] = p
		}
		index = allnum
	case "SectionType":
		switch op {
		case "-1":
			//获取全部博客
			p, allnum, _, err := models.QueryAllPostListByAjax(&u, page)
			if err == nil {
				this.Data["PostArr"] = p
			}
			index = allnum
		default:
			//获取版块博客
			p, allnum, _, err := models.QueryBySectionTypeAjax(&u, op, page)
			if err == nil {
				this.Data["PostArr"] = p
			}
			index = allnum
		}

	}

	this.SetSession("GategoryManagementType", types)
	this.SetSession("GategoryManagementOp", op)

	px := index / models.PageSize
	x := index % models.PageSize
	if x > 0 {
		px += 1
	}

	this.Data["Count"] = index
	//分页处理
	pl := PageLimit{CurrentIndex: page + 1, IndexCount: int(px)}
	pl.IndexArr = make([]int, int(px))
	for i := 0; i < pl.IndexCount; i++ {
		pl.IndexArr[i] = i + 1
	}
	this.Data["PageLimit"] = pl

	this.Data["OpType"]=utils.GenerateType(types,op)
	this.TplName = "user/profile-category.tpl"
}

/**
跳转到文章中心
分类
对外开放的方法，在Home控制器里使用到
*/
func Tocategory(this *beego.Controller, cate string, cond string, in int) {
	//导航栏
	arr := common.GetNavData(1, this)
	this.Data["NavArr"] = arr

	//分页下标
	var index int64
	var cnum int64
	if cate == "ContentType" {
		index, cnum = models.GetPageIndex(cate, cond, 0, true)
	} else {
		index, cnum = models.GetPageIndex(cate, cond, 0, false)
	}
	pl := PageLimit{CurrentIndex: in + 1, IndexCount: int(index)}
	pl.IndexArr = make([]int, int(index))

	for i := 0; i < pl.IndexCount; i++ {
		pl.IndexArr[i] = i + 1
	}
	this.Data["PageLimit"] = pl




	//博客内容
	if cate == "all" {
		//全部博客
		p, _, err := models.QueryPostList(in)
		if err == nil {
			this.Data["PostArr"] = p
		}
	} else if cate == "SectionType" {
		//版块博客
		p, _, err := models.QueryBySectionType(cond, in)
		if err == nil {
			this.Data["PostArr"] = p
		}
		this.SetSession("SectionType", cond) //设置正在操作的版块类型，0生活，1科技，2其他
	} else if cate == "Search" {
		//查找
		p, _, err := models.GetSearch(cond, in)
		if err == nil {
			this.Data["PostArr"] = p
		}
		this.Data["Search"] = cnum
	} else if cate == "ContentType" {
		//文章类型博客
		beego.Info("文章类型博客")

		p, _, err := models.QueryListByContentType(cond, in)
		if err == nil {
			this.Data["PostArr"] = p
		}
		this.SetSession("SectionType", cond) //设置正在操作的版块类型，0生活，1科技，2其他
	}


	//操作版块，版块分类（科技，生活，其他），全部文章，搜索,分类
	this.SetSession("Section", cate)

	InitSectionActive(this, cate, cond)
	//热门博客
	p1, _, err := models.GetHotBlog()
	if err == nil {
		this.Data["HotBlog"] = p1
	}

	life, technology, other, all := models.QuerySectionCount()
	this.Data["Life"] = life
	this.Data["Technology"] = technology
	this.Data["Other"] = other
	this.Data["All"] = all

	uname := this.GetSession("uname")
	if uname != nil {
		u := models.GetUserInfo(uname.(string))
		this.Data["CuerrentUser"] = u
	}
	this.TplName = "category.html"
}

/**
版块操作
*/
func (this *CategoryController) Section() {

	o := this.Input().Get("op")
	t := this.Input().Get("type")
	beego.Info("Section...操作类型", t, "操作方式", o)
	switch t {
	case "ContentType":
		Tocategory(&(this.Controller), "ContentType", o, 0)

	case "SectionType":
		if o == "-1" {
			Tocategory(&(this.Controller), "all", o, 0)
		} else {
			Tocategory(&(this.Controller), "SectionType", o, 0)
		}


	}

}

/**
分页操作，全部文章，版块，搜索内容，个人中心，管理员中心
*/

func (this *CategoryController) PageIndex() {
	op := this.Input().Get("op")
	form := this.Input().Get("form")
	s := this.Input().Get("index")
	var index int
	switch op {
	case "Index":
		i, _ := strconv.Atoi(s)
		index = i - 1
	case "Previous":
		i, _ := strconv.Atoi(s)
		index = i - 2
	case "Next":
		i, _ := strconv.Atoi(s)
		index = i
	}

	switch form {
	case "category":
		var cond string
		v := this.GetSession("SectionType")
		if v != nil {
			cond = v.(string)
		}
		cate := this.GetSession("Section").(string)

		Tocategory(&(this.Controller), cate, cond, index)
	case "usercenter":
		ToUserCenter(&(this.Controller), index)
	case "gategorymanage":
		t := this.GetSession("GategoryManagementType").(string)
		o := this.GetSession("GategoryManagementOp").(string)
		GategoryManagement(this, t, o, index)

	}

}

func (this *CategoryController) SetPollingPic() {
	id := this.Input().Get("pid")
	pos := this.Input().Get("posindex")
	//t := this.Input().Get("time")
	err := models.SetRollPic(id, pos)
	if err == nil {
		re := struct {
			Ok  int
			Msg string
		}{200, "Successs"}

		this.Data["json"] = &re
	} else {
		re := struct {
			Ok  int
			Msg string
		}{400, "Fail"}

		this.Data["json"] = &re
	}

	this.ServeJSON()

}

/**
搜索时，AJAX提示
*/
func (this *CategoryController) AjaxAutoTip() {
	b := this.Input().Get("term")

	data, num, err := models.AutoTip(b)
	reMsg := make([]string, num)
	if err == nil {
		for i, v := range data {
			reMsg[i] = v[0].(string)
		}
	}
	beego.Info("result", data, num, err)
	this.Data["json"] = &reMsg
	this.ServeJSON()
}

/**
查找搜索内容
*/
func (this *CategoryController) AjaxSearch() {
	v := this.Input().Get("autocomplete")
	Tocategory(&(this.Controller), "Search", v, 0)
}

/**
浏览单篇文章
*/
func (this *CategoryController) BrowseCategory() {
	id := this.Input().Get("id")
	p, err := models.GetOne("Id", id)
	if err == nil {
		this.Data["Title"] = p.Title
		this.Data["Post"] = p
		this.Data["fali"] = false

	} else {
		this.Data["Title"] = "加载失败"
		this.Data["Content"] = "<h1>文章已下线或删除</h1>"
		this.Data["fali"] = true
	}
	arr := common.GetNavData(-1, &(this.Controller))
	this.Data["NavArr"] = arr
	this.TplName = "visibility.html"
}

/**
初始化版块激活状态
*/
func InitSectionActive(this *beego.Controller, op string, cond string) {

	if op == "all" {
		this.Data["AllActive"] = true
	} else if op == "SectionType" {
		switch cond {
		case "0":
			this.Data["LifeActive"] = true
		case "1":
			this.Data["TechnologyActive"] = true
		case "2":
			this.Data["OtherActive"] = true
		}
	} else if op == "Search" {
		this.Data["SearchActive"] = true
	}
}

func (this *CategoryController) Prepare() {
	this.EnableXSRF = false
	CheckUserLogin(&(this.Controller))
}
