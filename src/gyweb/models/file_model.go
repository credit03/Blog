package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

func ImageTempFileAdd(path string) error {
	i := &ImageTemp{Url:path}
	o := orm.NewOrm()
	_, err := o.Insert(i)
	return err

}

func VideoTempFileAdd(path string) error {
	i := &VideoTemp{Url:path}
	o := orm.NewOrm()
	_, err := o.Insert(i)
	return err

}

func GetOneImageById(pid int64) string {
	var p Post
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("ContentType", 1).Or("ContentType", 2)
	err := o.QueryTable("post").Filter("Id", pid).SetCond(cond1).One(&p)
	if err == nil {
		o.LoadRelated(&p, "Images")
		return p.Images[0].Url
	}

	return "/static/img/2.jpg"

}
func GetImagePosts(page int, max int) ([]Post, int64, error) {
	return GetImagePostsByName("-all", "", page, max)
}

func GetImages(uname string, page int) []Post {
	p, _, _ := GetImagePostsByName(uname, "", page, 8)

	return p
}

func GetRollPic() []RollPic {
	var r []RollPic
	o := orm.NewOrm()
	o.QueryTable("RollPic").OrderBy("Id").All(&r)
	return r
}

func SetRollPic(id string, posid string) error {
	rid, _ := strconv.Atoi(posid)
	o := orm.NewOrm()
	p, err := GetOnebyColumn("Id", id)
	if err == nil {
		r := &RollPic{Id:int64(rid), BlongPostId:p.Id, ImageUrl:p.Images[0].Url, Title:p.Title, Describe:p.Describe}
		_, err = o.Update(r)
	}
	return err

}

func GetImagePostsByName(uname string, order string, page int, MAX int) ([]Post, int64, error) {

	if page == 0 {
		page = 1
	}
	if len(order) == 0 {
		order = "-Created"
	}
	var p []Post
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("ContentType", 1).Or("ContentType", 2)
	var num int64
	var err error
	switch uname {
	case "-all":
		num, err = o.QueryTable("post").OrderBy(order).SetCond(cond1).Limit(page * MAX).RelatedSel().All(&p)
	default:
		cond2 := cond.AndCond(cond.And("Author__Name", uname)).AndCond(cond1)
		num, err = o.QueryTable("post").OrderBy(order).SetCond(cond2).Limit(page * MAX).RelatedSel().All(&p)


	}

	if err == nil {
		for i, _ := range p {
			o.LoadRelated(&p[i], "Images")
		}
	}

	return p, num, err
}

func GetImageCountPage() (int64) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("ContentType", 1).Or("ContentType", 2)
	i, _ := o.QueryTable("post").SetCond(cond1).Count()
	p := i / 8
	x := i % 8
	if x > 0 {
		p += 1
	}
	return p

}