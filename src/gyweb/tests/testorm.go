package main

import (
	"../models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

/*func init() {
	models.RegisterDB()
}*/
func main() {
	fmt.Println(time.Now())
}

func TestOrm()  {
	/*orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", true, true)

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	u := &models.User{Name:"cre", Pic:"www.baidu.com", Pwd:"123456"}

	p1 := &models.Post{Title:"test1", Content:"如果需要通过条件", Author:u, Created:time.Now(), Updated:time.Now(), ReplyTime:time.Now()}
	p2 := &models.Post{Title:"test2", Content:"1234567890", Author:u, Created:time.Now(), Updated:time.Now(), ReplyTime:time.Now()}

	c1 := &models.Comment{Content:"d1", Uid:1, Time:time.Now(), Belong:p1}
	time.Sleep(1 * time.Second)
	c2 := &models.Comment{Content:"d2", Uid:1, Time:time.Now(), Belong:p1}
	time.Sleep(1 * time.Second)
	c3 := &models.Comment{Content:"d3", Uid:1, Time:time.Now(), Belong:p2}
	time.Sleep(1 * time.Second)
	c4 := &models.Comment{Content:"d4", Uid:1, Time:time.Now(), Belong:p2}

	fmt.Println(o.Insert(u))

	fmt.Println(o.Insert(p1))
	fmt.Println(o.Insert(p2))

	fmt.Println(o.Insert(c1))
	fmt.Println(o.Insert(c2))
	fmt.Println(o.Insert(c3))
	fmt.Println(o.Insert(c4))

	fmt.Println("直接查询。。。。")
	var posts []*models.Post
	num, err := o.QueryTable("post").Filter("Author", 1).RelatedSel().All(&posts)
	if err == nil {
		fmt.Printf("%d posts read\n", num)
		for _, post := range posts {
			fmt.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.Author.Name, post.Title)
		}
	}

	var cos []*models.Comment
	num, err = o.QueryTable("comment").Filter("Belong", 1).RelatedSel().All(&cos)

	if err == nil {
		fmt.Printf("%d comment read\n", num)
		for _, con := range cos {
			fmt.Printf("Id: %d, UserName: %s, CONTENT: %s; 帖子内容：%s\n", con.Id, con.Belong.Author.Name, con.Content, con.Belong.Content)
		}
	}

	//根据 User表的Posts.Author.Id 查询对应的 User：

	var u1 []*models.User
	num, err = o.QueryTable("user").Filter("Posts__Author__Id", "1").Limit(10).All(&u1)
	if err == nil {

		fmt.Printf("%d 查询对应的 User\n", num)
		for _, con := range u1 {
			fmt.Printf("Id: %d; UserName: %s; pic:%s\n", con.Id, con.Name, con.Pic)
		}
	}*/
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	p1, _, _ := Query(0)
	//p2, _, _ := Query(0)
	for _, v := range p1 {

		fmt.Println("Tile:", v.Title)
		fmt.Println("Content:", v.Content)
		fmt.Println("ContentType:", v.ContentType)
		fmt.Println("SectionType:", v.SectionType)
		fmt.Println("Describe", v.Describe)
		fmt.Println("Author", v.Author)

		if v.ContentType == 1 || v.ContentType == 2 {
			for _, v1 := range v.Images {
				fmt.Println("图片链接", v1.Url)
			}
		}
		fmt.Println("VideoUrl", v.VideoUrl)
		fmt.Println("**************************************")
	}
	//	fmt.Println("标题", p[0].Title, p[0].Images[0].Url, "大小", num, err)
	GetPageIndex()
}
func Query(page int) ([]models.Post, int64, error) {

	var p []models.Post
	o := orm.NewOrm()
	num, err := o.QueryTable("post").OrderBy("-Created").Limit(10, 10 * page).RelatedSel().All(&p)
	fmt.Println("num", num, "err:", err)
	for i, v := range p {
		if v.ContentType == 1 || v.ContentType == 2 {
			o.LoadRelated(&p[i], "Images")
		}
	}
	return p, num, err
}

func GetPageIndex() {
	i := QueryPostCount()
	p := i / 10
	x := i % 10
	fmt.Println(i, p, x)

}
func QueryPostCount() int64 {
	o := orm.NewOrm()
	num, _ := o.QueryTable("post").Count()
	return num

}
