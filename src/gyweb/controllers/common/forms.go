package common

type UserPost struct {
	Uname     string `form:"uname"`
	Pwd       string `form:"pwd"`
	Autologin string  `form:"autologin"`
}

type UpdateUserForm struct {
	NickName string `form:"nickname"`
	Sex      int    `form:"sex"`
	Age      int    `form:"age"`
	Birth    string `form:"birth"`
	Address  string `form:"address"`
}

type PublishForm struct {
	Id          string  `form:"pid"`
	Title       string `form:"title"`
	Content     string `form:"content"`
	ContentType string `form:"contenttype"`
	Describe    string `form:"describe"`
	Sectiontype string `form:"sectiontype"`
	Useimage    string `form:"useimage"`
	Nouseimage  string `form:"nouseimage"`
	ImageLabel  string `form:"imagelabel"`

	/**
	   下面两个字段是备用
	 */
	Usevideo    string `form:"usevideo"`
	Nousevideo  string `form:"nousevideo"`
}
