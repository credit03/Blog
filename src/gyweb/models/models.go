package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	_"github.com/mattn/go-sqlite3"
	"github.com/Unknwon/com"
	"os"
	"path"
)

// 数据库建表总models
//
const (
	_DB_NAME = "data/beeblog.db"
	_MYSQL_DRIVER = "mysql"
	_SQLITE3_DRIVER = "sqlite3"
)

/**
用户
 */
type User struct {
	Id       int64                          //int, int32 - 设置 auto 或者名称为 Id 时 integer AUTO_INCREMENT
	Name     string `orm:"unique;size(24)"` //设置不为空并不能大小24个字符
	NickName string `orm:"unique;size(24)"` //昵称
	Sex      bool   `orm:"default(false)"`  //性别
	Age      int    `orm:"default(0)"`
	Birth    string  `orm:"default(2016-10-10)"`
	Address  string `orm:"default(广州市天河区)"`
	Pwd      string `orm:"unique;size(64)"` //设置不为空并不能大小24个字符
	Pic      string
	IsAdmin  bool   `orm:"default(false)"`
	Posts    []*Post `orm:"reverse(many)"`  // 设置一对多的反向关系
}
/**
帖子
 */
type Post struct {
	Id              int64
	ContentType     int64                             //0:表示纯文本文章，1代表单张图片文章，2代表多张图片文章，3代表视频文章
	Title           string `orm:"unique"`
	Content         string `orm:"size(20480)"`
	SectionType     int                               //版块0,1,2
	Describe        string `orm:"size(512)"`          //描述
	Author          *User  `orm:"rel(fk)"`            //设置一对多关系(外键)
	Created         time.Time `orm:"index"`           //为单个字段增加索引，例如用户选择某天xx时的文章，就可以根据这个索引匹配
	Updated         time.Time `orm:"index"`
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	BrowseCount     int64
	Comments        []*Comment  `orm:"reverse(many)"` // 设置一对多的反向关系
	Images          []*Image  `orm:"reverse(many)"`   // 设置一对多的反向关系
	ImageLabel      string
	VideoUrl        string
							  //Video           *Video `orm:"null;rel(one);on_delete(set_null)"` //RelOneToOne:
}

type RollPic struct {
	Id          int64                    //轮播图位置
	Title       string `orm:"unique"`
	Describe    string `orm:"size(256)"` //描述
	ImageUrl    string `orm:"unique"`
	BlongPostId int64
}

/**
图片
 */
type Image struct {
	Id     int64
	Url    string `orm:"unique"`
	Belong *Post `orm:"rel(fk)"`
}
/**
视频
 */
type Video struct {
	Id  int64
	Url string `orm:"unique"`
	//Belong *Post `orm:"reverse(one)"`
}

/**
图片临时表
 */
type ImageTemp struct {
	Id  int64
	Url string `orm:"unique"`
}
/**
视频临时表
 */
type VideoTemp struct {
	Id  int64
	Url string `orm:"unique"`
}
/**
评论
 */
type Comment struct {
	Id      int64
	Uid     int64
	Time    time.Time
	Content string`orm:"size(200)"`
	Belong  *Post `orm:"rel(fk)"`
}

/**
注册
 */
func RegisterDB() {
	/**
	Using

	切换为其他数据库

	orm.RegisterDataBase("db1", "mysql", "root:root@/orm_db2?charset=utf8")
	orm.RegisterDataBase("db2", "sqlite3", "data.db")

	o1 := orm.NewOrm()
	o1.Using("db1")

	o2 := orm.NewOrm()
	o2.Using("db2")

	// 切换为其他数据库以后
	// 这个 Ormer 对象的其下的 api 调用都将使用这个数据库


	默认使用 default 数据库，无需调用 Using
	 */

	//使用多个数据库
	orm.RegisterModel(new(Comment), new(Post), new(User), new(Image), new(Video), new(ImageTemp), new(VideoTemp), new(RollPic))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	//设置默认数据库
	//Go语言使用Beego的ORM插入Mysql后，时区不一致的解决方案&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:java@/my_db?charset=utf8&loc=Asia%2FShanghai", 50)


	//判断是否创建了sqlite3数据库
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("db2", _SQLITE3_DRIVER, _DB_NAME, 20)
}

func AutoCreateAdmin() {

	o := orm.NewOrm()
	b := o.QueryTable("user").Filter("Name", "admin").Exist()
	if !b {
		u := &User{Name:"admin", Pwd:"admin", Pic:"/static/img/1.jpg"}
		o.Insert(u)

	}

}

func AutoCreateRollPic() {
	o := orm.NewOrm()
	c, _ := o.QueryTable("RollPic").Count()
	if c == 0 {

		p1 := &RollPic{Id:1, ImageUrl: "/static/img/back-1.jpg", Title: "经济发展不是“数学题” 坐而论道不如聚力改革", Describe: "　9月3日，国家主席习近平在杭州出席2016年二十国集团工商峰会开幕式，并发表题为《中国发展新起点 全球增长新蓝图》的主旨演讲。演讲中，他回应回应外界对中国经济质疑：“行胜于言。中国用实际行动对这些问题作出了回答。”", BlongPostId:4}
		p2 := &RollPic{Id:2, ImageUrl: "/static/img/back-2.jpg", Title: "Go 机器学习库 Gorgonia", Describe: "Gorgonia 是 Go 机器学习库。撰写和评估多维数组的数学公式。", BlongPostId:4}
		p3 := &RollPic{Id:3, ImageUrl: "/static/img/back-3.jpg", Title: "Pig and Dan.", Describe:"My Life", BlongPostId:4}
		o.Insert(p1)
		o.Insert(p2)
		o.Insert(p3)
	}
}