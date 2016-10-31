package common

import (
	"fmt"
	"math/rand"

	"github.com/astaxie/beego"
)

/**
导航栏ITEM
*/
type Nav struct {
	Active bool
	Url    string
	Name   string
	IsShow bool
}

/**
滚动图片
*/
type RollingPicture struct {
	Active   bool
	ImageUrl string
	Title    string
	Content  string
}

type HotInfo struct {
	ImageUrl string
	Title    string
	Content  string
	Url      string
}

func init() {
	initNav()
	//initRollingPicture()
}

var (
	navdata            []Nav
	rollingPicturedata []RollingPicture
)

const (
	//用户操作
	To_login_Page = "/user/loginpage"
	To_login = "/user/login"
	To_register = "/user/register"
	To_userinfo_browse = "/user/borwseinfo"
	To_userinfo_edit = "/myajax/editinfo/:id"
	To_userinfo_update = "/myajax/updateinfo/:id"

	To_userpic_edit = "/myajax/eidtuserpic/:id"
	To_userpic_update = "/myajax/updateuserpic/:id"

	To_userpwd_edit = "/myajax/eidtuserpwd"
	To_userpwd_update = "/myajax/updateuserpwd"
	//文章操作
	To_Category = "/?op=category"
	To_Publish = "/myajax/publish"
	To_Category_Add = "/myajax/publish/toadd"
	To_Category_Delete = "/myajax/category/delete"
	To_Category_Edit = "/myajax/category/edit"
	To_Category_Updtae = "/myajax/category/update"
	To_Category_Browse = "/browse_category"

	//管理员后台操作
	To_User_Center = "/myajax/usercenter"
	To_Center_Gategory_manage = "/myajax/gategory_manage/:id"
	To_User_Manage = "/myajax/user_manage/:id"

	To_RollingPic_Set = "/myajax/setrollingpicture"

	//文章ajax搜索提示
	AutoTip = "/ajaxautotip"

	//文件操作
	Image_Upload = "/iamgeupload"
	Viedo_Upload = "/videoupload"
)

func initNav() {
	home := Nav{Url: "#", Name: "首页", IsShow: true}
	category := Nav{Url: To_Category, Name: "文章", IsShow: true}
	//publish := Nav{Url:"/?op=publish", Name:"发表博客"}
	publish := Nav{Url: To_Publish, Name: "发表", IsShow: true}
	gallery := Nav{Url: "/?op=gallery", Name: "图库", IsShow: true}

	usercenter := Nav{Url: To_User_Center, Name: "个人中心", IsShow: false}
	//admincenter := Nav{Url:To_Admin_Center, Name:"控制台", IsShow:false}

	about := Nav{Url: "/?op=about", Name: "关于Cre", IsShow: true}
	//	navdata = []Nav{home, category, publish, gallery, usercenter, admincenter, about}
	navdata = []Nav{home, category, publish, gallery, usercenter, about}
}
func initRollingPicture() {
	p1 := RollingPicture{ImageUrl: "/static/img/back-1.jpg", Title: "经济发展不是“数学题” 坐而论道不如聚力改革", Content: "　9月3日，国家主席习近平在杭州出席2016年二十国集团工商峰会开幕式，并发表题为《中国发展新起点 全球增长新蓝图》的主旨演讲。演讲中，他回应回应外界对中国经济质疑：“行胜于言。中国用实际行动对这些问题作出了回答。”"}
	p2 := RollingPicture{ImageUrl: "/static/img/back-2.jpg", Title: "Go 机器学习库 Gorgonia", Content: "Gorgonia 是 Go 机器学习库。撰写和评估多维数组的数学公式。与  [Theano](http://deeplearning.net/software/theano/) 和 [TensorFlow](https://www.tensorflow.org/) 理念相似。具体来说， #### Gorgonia 性能： * 执行自动分化 * 执行符号微分 * 优化 梯度下降 * 进行稳定的数值计算 * 提供便捷功能来帮助创建神经网络 * 操作快（与Theano和Tensorflow速度相当） * 支持GPU / CUDA * 支持分布式计算 "}
	p3 := RollingPicture{ImageUrl: "/static/img/back-3.jpg", Title: "Pig and Dan.", Content: "My Life"}
	rollingPicturedata = []RollingPicture{p1, p2, p3}
}

func GetNavData(op int, this *beego.Controller) []Nav {

	if op != -1 {
		navdata[0].Url = "/"
	}
	/**
	注意一定要清空这两个参数，全局变量，整个网站都生效的
	*/
	navdata[4].IsShow = false
	navdata[4].Name = "个人中心"
	//navdata[5].IsShow = false

	isadmin := this.GetSession("isadmin")
	beego.Info("是否是管理员", isadmin)
	if isadmin != nil {
		navdata[4].IsShow = true
		b := isadmin.(bool)
		if b == true {
			navdata[4].Name = "控制台"
		} else {
			navdata[4].Name = "个人中心"
		}
	}
	/**
	  清空选择
	*/
	for k, _ := range navdata {
		navdata[k].Active = false
	}
	if op != -1 && op != -2 {
		navdata[op].Active = true
	}
	return navdata
}

func GetRollingPictureData(op int) []RollingPicture {

	/**
	  清空选择
	*/
	for k, _ := range rollingPicturedata {
		rollingPicturedata[k].Active = false
	}
	rollingPicturedata[op].Active = true
	return rollingPicturedata
}

func GetHotInfo() []HotInfo {
	index := rand.Intn(32)
	h1 := HotInfo{ImageUrl: fmt.Sprint("/static/img/", index, ".jpg"), Title: "Dan", Content: "StructTag 的定义用的标签用为form，和 ParseForm 方法 共用一个标签，标签后面有三个可选参数，用,分割。第一个参数为表单中类型的name的值，如果为空，则以struct field name为值。第二个参数为表单组件的类型，如果为空，则为text。表单组件的标签默认为struct field name的值，否则为第三个值。", Url: "#"}
	index = rand.Intn(32)
	h2 := HotInfo{ImageUrl: fmt.Sprint("/static/img/", index, ".jpg"), Title: "Gy", Content: "如果form标签只有一个值，则为表单中类型name的值，除了最后一个值可以忽略外，其他位置的必须要有,号分割，如：form:,,姓名：", Url: "#"}
	index = rand.Intn(32)
	h3 := HotInfo{ImageUrl: fmt.Sprint("/static/img/", index, ".jpg"), Title: "Dan and Gy", Content: "如果form标签只有一个值，则为表单中类型name的值，除了最后一个值可以忽略外，其他位置的必须要有,号分割，如：form:,,姓名：", Url: "#"}
	index = rand.Intn(32)
	h4 := HotInfo{ImageUrl: fmt.Sprint("/static/img/", index, ".jpg"), Title: "Dan and Gy", Content: "如果form标签只有一个值，则为表单中类型name的值，除了最后一个值可以忽略外，其他位置的必须要有,号分割，如：form:,,姓名：", Url: "#"}
	hotinfodata := []HotInfo{h1, h2, h3, h4}
	return hotinfodata
}
