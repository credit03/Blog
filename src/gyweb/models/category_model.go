package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"../controllers/common"
	"../utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	PageSize = 5
)

func CategoryAdd(f *common.PublishForm, uname string) (int64, error) {

	o := orm.NewOrm()
	err := o.Begin() //开启事务
	u := &User{Name: uname}
	//Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	err = o.Read(u, "Name")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		t := time.Now()
		beego.Info("添加文章时间", t.Format("2006-01-02 15:04:05"))
		ct, _ := strconv.Atoi(f.ContentType)
		se, _ := strconv.Atoi(f.Sectiontype)
		p := &Post{Title: f.Title, Content: f.Content, ContentType: int64(ct), SectionType: se, Describe: f.Describe, Author: u, Created: t, Updated: t, ReplyTime: t, ImageLabel: f.ImageLabel}
		if len(f.Usevideo) > 0 {
			p.VideoUrl = f.Usevideo //设置视频链接
		}
		var id int64
		id, err = o.Insert(p) //插入帖子

		if ct == 1 || ct == 2 {
			sub := strings.Split(f.Useimage, "#substr#")
			sub1 := strings.Split(f.Nouseimage, "#substr#")
			l := len(sub)
			if sub != nil && l > 0 {
				for i := 0; i < l - 1; i++ {
					v := utils.Substr(sub[i], 1, len(sub[i]))
					it := &ImageTemp{Url: v}
					err1 := o.Read(it, "url")
					if err1 == nil {
						img := &Image{Url: it.Url, Belong: p}
						o.Insert(img)
						o.Delete(it)
					}
				}
			}

			l = len(sub1)
			if sub1 != nil && len(sub1) > 0 {
				for i := 0; i < l - 1; i++ {
					v := utils.Substr(sub1[i], 1, len(sub1[i]))
					it := &ImageTemp{Url: v}
					o.Delete(it)
					err2 := os.Remove(v)
					if err2 != nil {
						//如果删除失败则输出 file remove Error!
						fmt.Println("file remove Error!")
						//输出错误详细信息
						fmt.Printf("%s", err2)
					} else {
						//如果删除成功则输出 file remove OK!
						fmt.Println("如果删除成功则输出 file remove OK!")
					}

				}
			}

		}

		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
			return id, err
		}

	}
	return -1, errors.New("插入失败")
}

func QueryPostList(page int) ([]Post, int64, error) {

	var p []Post
	o := orm.NewOrm()
	//                     Limit  限制最大返回数据行数，第二个参数可以设置 Offset(偏移);RelatedSel 关系查询;All 返回对应的结果集对象
	num, err := o.QueryTable("post").OrderBy("-Created").Limit(PageSize, PageSize * page).RelatedSel().All(&p)
	for i, v := range p {
		if v.ContentType == 1 || v.ContentType == 2 {
			/**
			载入关系字段
			LoadRelated 用于载入模型的关系字段，包括所有的 rel/reverse - one/many 关系
			ManyToMany 关系字段载入
			*/
			o.LoadRelated(&p[i], "Images")
		}

	}
	return p, num, err
}

func QueryAllPostListByAjax(u *User, page int) ([]Post, int64, int64, error) {

	var p []Post
	o := orm.NewOrm()
	var num, allnum int64
	var err error

	if u.IsAdmin {
		num, err = o.QueryTable("post").OrderBy("-Created").Limit(PageSize, PageSize * page).RelatedSel().All(&p)
		allnum, _ = o.QueryTable("post").Count()
	} else {
		num, err = o.QueryTable("post").Filter("Author", u).OrderBy("-Created").Limit(PageSize, PageSize * page).RelatedSel().All(&p)
		allnum, _ = o.QueryTable("post").Filter("Author", u).Count()
	}
	for i, v := range p {
		if v.ContentType == 1 || v.ContentType == 2 {
			/**
			载入关系字段
			LoadRelated 用于载入模型的关系字段，包括所有的 rel/reverse - one/many 关系
			ManyToMany 关系字段载入
			*/
			o.LoadRelated(&p[i], "Images")
		}

	}
	return p, allnum, num, err
}

func QueryBySectionTypeAjax(u *User, section string, page int) ([]Post, int64, int64, error) {
	var p []Post
	o := orm.NewOrm()
	v, _ := strconv.Atoi(section)
	var allnum, num int64
	var err error
	var qs orm.QuerySeter
	if u.IsAdmin {
		qs = o.QueryTable("post").Filter("SectionType", v)
	} else {
		qs = o.QueryTable("post").Filter("Author", u).Filter("SectionType", v)

	}

	allnum, _ = qs.Count()
	num, err = qs.OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)

	for i, v := range p {
		if v.ContentType == 1 || v.ContentType == 2 {
			o.LoadRelated(&p[i], "Images")
		}
	}

	return p, allnum, num, err
}
func QueryByContentTypeAjax(u *User, section string, page int) ([]Post, int64, int64, error) {
	var p []Post
	o := orm.NewOrm()
	v, _ := strconv.Atoi(section)

	var num, allnum int64
	var err error
	if v != 1 {
		if v == 2 {
			v = 3
		}
		var qs orm.QuerySeter
		if u.IsAdmin {
			qs = o.QueryTable("post").Filter("ContentType", v)
		} else {
			qs = o.QueryTable("post").Filter("Author", u).Filter("ContentType", v)

		}
		allnum, _ = qs.Count()
		num, err = qs.OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)
	} else {

		cond := orm.NewCondition()
		cond1 := cond.And("ContentType", 1).Or("ContentType", 2)
		var qs orm.QuerySeter
		if u.IsAdmin {
			qs = o.QueryTable("post").SetCond(cond1)
		} else {
			cond2 := cond.AndCond(cond.And("Author", u)).AndCond(cond1)
			qs = o.QueryTable("post").SetCond(cond2)

		}

		allnum, _ = qs.Count()
		num, err = qs.OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)
		for i, _ := range p {
			o.LoadRelated(&p[i], "Images")

		}
	}

	return p, allnum, num, err
}
func GetHotBlog() ([]Post, int64, error) {
	/*var p []Post
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("BrowseCount__gt", 5).Or("ReplyCount__gt", 5)
	num, err := o.QueryTable("post").OrderBy("-BrowseCount").SetCond(cond1).Limit(2).RelatedSel().All(&p)
	if err != nil || num == 0 {
		num, err = o.QueryTable("post").OrderBy("-BrowseCount").Limit(2).RelatedSel().All(&p)
	}*/
	return GetHotBlogByName("", 2)
}
func GetHotBlogByName(uname string, count int) ([]Post, int64, error) {
	var p []Post
	var num int64
	var err error
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("BrowseCount__gt", 5).Or("ReplyCount__gt", 5)

	if utils.StringIsNotEmpty(uname) {
		u := &User{Name: uname}
		o.Read(u, "Name")
		cond2 := cond.AndCond(cond.And("Author", u)).AndCond(cond1)
		num, err = o.QueryTable("post").SetCond(cond2).OrderBy("-BrowseCount").Limit(count).RelatedSel().All(&p)
		if err != nil || num == 0 {
			num, err = o.QueryTable("post").Filter("Author", u).OrderBy("-BrowseCount").Limit(count).RelatedSel().All(&p)
		}
	} else {

		num, err = o.QueryTable("post").OrderBy("-BrowseCount").SetCond(cond1).Limit(count).RelatedSel().All(&p)
		if err != nil || num == 0 {
			num, err = o.QueryTable("post").OrderBy("-BrowseCount").Limit(count).RelatedSel().All(&p)
		}
	}
	return p, num, err
}

func GetGategoryCountAndBrowseCount(uname string) (int64, int64) {
	var p []Post
	var bc int64 = 0
	o := orm.NewOrm()
	num, err := o.QueryTable("post").Filter("Author__Name", uname).All(&p)
	if err == nil {
		for _, v := range p {
			bc += v.BrowseCount
		}
	}

	return num, bc
}

func AutoTip(query string) ([]orm.ParamsList, int64, error) {
	var lists []orm.ParamsList
	o := orm.NewOrm()
	num, err := o.QueryTable("post").Filter("Title__icontains", query).Distinct().Limit(10).ValuesList(&lists, "Title")
	return lists, num, err
}

func GetSearch(search string, page int) ([]Post, int64, error) {
	var p []Post
	o := orm.NewOrm()
	num, err := o.QueryTable("post").Filter("Title__icontains", search).OrderBy("-BrowseCount").Limit(PageSize, PageSize * page).RelatedSel().All(&p)
	return p, num, err
}

func GetPageIndex(filed string, val string, Size int64, IsContentType bool) (int64, int64) {

	var i int64
	if IsContentType {
		i = QueryContentTypeCount(filed, val)
	} else {
		i = QueryPostCount(filed, val)
	}
	var p int64
	if Size == 0 {
		p = i / PageSize
		x := i % PageSize
		if x > 0 {
			p += 1
		}
	} else {
		p = i / Size
		x := i % Size
		if x > 0 {
			p += 1
		}
	}

	return p, i

}
func QueryPostCount(filed string, val string) int64 {
	o := orm.NewOrm()
	if filed == "all" {
		num, _ := o.QueryTable("post").Count()
		return num
	} else if filed == "SectionType" {
		v, _ := strconv.Atoi(val)
		num, _ := o.QueryTable("post").Filter("SectionType", v).Count()
		return num
	} else if filed == "Search" {
		num, _ := o.QueryTable("post").Filter("Title__icontains", val).Count()
		return num
	} else {

		num, _ := o.QueryTable("post").Filter(filed, val).Count()
		return num
	}

	return 0
}
func QueryContentTypeCount(filed string, val string) int64 {
	o := orm.NewOrm()
	v, _ := strconv.Atoi(val)
	var num int64
	if v != 1 {
		if v == 2 {
			v = 3
		}
		num, _ = o.QueryTable("post").Filter("ContentType", v).Count()
	} else {

		cond := orm.NewCondition()
		cond1 := cond.And("ContentType", 1).Or("ContentType", 2)
		num, _ = o.QueryTable("post").SetCond(cond1).Count()

	}
	return num

}
func QueryListByContentType(section string, page int) ([]Post, int64, error) {
	var p []Post
	o := orm.NewOrm()
	v, _ := strconv.Atoi(section)

	var num int64
	var err error
	if v != 1 {
		if v == 2 {
			v = 3
		}
		num, err = o.QueryTable("post").Filter("ContentType", v).OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)
	} else {

		cond := orm.NewCondition()
		cond1 := cond.And("ContentType", 1).Or("ContentType", 2)
		num, err = o.QueryTable("post").SetCond(cond1).OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)
		for i, _ := range p {
			o.LoadRelated(&p[i], "Images")

		}
	}

	return p, num, err
}
func QueryBySectionType(section string, page int) ([]Post, int64, error) {
	var p []Post
	o := orm.NewOrm()
	v, _ := strconv.Atoi(section)
	num, err := o.QueryTable("post").Filter("SectionType", v).OrderBy("-Created").Limit(PageSize, page * PageSize).RelatedSel().All(&p)
	for i, v := range p {
		if v.ContentType == 1 || v.ContentType == 2 {
			o.LoadRelated(&p[i], "Images")
		}
	}
	return p, num, err
}

func QuerySectionCount() (int64, int64, int64, int64) {

	o := orm.NewOrm()
	one, _ := o.QueryTable("post").Filter("SectionType", 0).Count()
	two, _ := o.QueryTable("post").Filter("SectionType", 1).Count()
	three, _ := o.QueryTable("post").Filter("SectionType", 2).Count()
	all, _ := o.QueryTable("post").Count()
	return one, two, three, all
}

func GetOne(filedtype string, value string) (Post, error) {
	return GetOneByUserName(filedtype, value, "", true)
}



func GetOneByUserName(filedtype string, value string, uname string, coladd bool) (Post, error) {
	p := new(Post)
	o := orm.NewOrm()

	var err error
	if coladd {
		//原子操作增加字段值
		_, err = o.QueryTable("post").Filter(filedtype, value).Update(orm.Params{"BrowseCount": orm.ColValue(orm.ColAdd, 1)})
	}

	if utils.StringIsNotEmpty(uname) {
		u := &User{Name:uname}
		o.Read(u, "Name")
		if u.IsAdmin {
			err = o.QueryTable("post").Filter(filedtype, value).RelatedSel().One(p)
		} else {
			err = o.QueryTable("post").Filter(filedtype, value).Filter("Author", u).RelatedSel().One(p)
		}
	} else {
		err = o.QueryTable("post").Filter(filedtype, value).RelatedSel().One(p)

	}

	if err == nil {
		o.LoadRelated(p, "Images")
		o.LoadRelated(p, "Comments")
	}
	return *p, err
}

func GetOnebyColumn(filedtype string, value string) (Post, error) {
	p := new(Post)
	o := orm.NewOrm()
	err := o.QueryTable("post").Filter(filedtype, value).One(p)
	if err == nil {
		o.LoadRelated(p, "Images")
	}
	return *p, err
}

func QueryListByUserName(uname string, page int) ([]Post, int64, int64, error) {

	u := &User{Name: uname}
	o := orm.NewOrm()
	err := o.Read(u, "Name")
	if err == nil {
		i, num := GetPageIndex("Author", fmt.Sprint(u.Id), 15, false)
		q, _, err := QueryListByColumn("Author", fmt.Sprint(u.Id), "-Created", 15, page)
		return q, i, num, err
	} else {
		return []Post{}, 0, 0, errors.New("查找不到")
	}

}

func QueryListByColumn(fied string, val string, orderBy string, limitMax int, pageIndex int) ([]Post, int64, error) {

	o := orm.NewOrm()
	var p []Post
	num, err := o.QueryTable("post").Filter(fied, val).OrderBy(orderBy).Limit(limitMax, limitMax * pageIndex).RelatedSel().All(&p)
	return p, num, err

}

func DeletePostByID(id string, uname string) error {
	o := orm.NewOrm()
	u := &User{Name:uname}
	o.Read(u, "Name")
	var err error
	if u.IsAdmin {
		_, err = o.QueryTable("post").Filter("Id", id).Filter("Author__Name", uname).Delete()
	} else {
		_, err = o.QueryTable("post").Filter("Id", id).Filter("Author", u).Delete()
	}
	return err
}

func UdpateCateGory(f *common.PublishForm, uname string) (int64, error) {

	o := orm.NewOrm()
	err := o.Begin() //开启事务
	u := &User{Name:uname}
	o.Read(u, "Name")
	var b bool
	if u.IsAdmin {
		b = o.QueryTable("post").Filter("Id", f.Id).Exist()
	} else {
		b = o.QueryTable("post").Filter("Id", f.Id).Filter("Author", u).Exist()

	}
	if b {
		t := time.Now()
		pid, _ := strconv.Atoi(f.Id)
		ct, _ := strconv.Atoi(f.ContentType)
		se, _ := strconv.Atoi(f.Sectiontype)
		p := &Post{Id: int64(pid), Title: f.Title, Content: f.Content, ContentType: int64(ct), SectionType: se, Describe: f.Describe, Updated: t, ImageLabel: f.ImageLabel}
		if len(f.Usevideo) > 0 {
			p.VideoUrl = f.Usevideo //设置视频链接
		}
		var id int64
		id, err = o.Update(p, "Title", "Content", "ContentType", "SectionType", "Describe", "Updated", "ImageLabel") //插入帖子
		beego.Info("UdpateCateGory 更新", err)
		if ct == 1 || ct == 2 {
			sub := strings.Split(f.Useimage, "#substr#")
			sub1 := strings.Split(f.Nouseimage, "#substr#")
			l := len(sub)
			if sub != nil && l > 0 {

				for i := 0; i < l - 1; i++ {

					v := utils.Substr(sub[i], 1, len(sub[i]))
					if !o.QueryTable("image").Filter("Url", v).Exist() {
						it := &ImageTemp{Url: v}
						err1 := o.Read(it, "url")
						if err1 == nil {
							img := &Image{Url: it.Url, Belong: p}
							o.Insert(img)
							o.Delete(it)
						}

					}

				}
			}

			l = len(sub1)
			if sub1 != nil && len(sub1) > 0 {
				for i := 0; i < l - 1; i++ {
					v := utils.Substr(sub1[i], 1, len(sub1[i]))
					it := &ImageTemp{Url: v}
					ig := &Image{Url: v}
					o.Delete(it, "Url")
					o.Delete(ig, "Url")
					err2 := os.Remove(v)
					if err2 != nil {
						//如果删除失败则输出 file remove Error!
						fmt.Println("file remove Error!")
						//输出错误详细信息
						fmt.Printf("%s", err2)
					} else {
						//如果删除成功则输出 file remove OK!
						fmt.Println("如果删除成功则输出 file remove OK!")
					}

				}
			}

		}

		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
			return id, err
		}

	}
	return -1, errors.New("更新失败")

}
