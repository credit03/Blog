package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"../controllers/common"
)

func CheckUser(uname string) bool {

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	con := qs.Filter("Name", uname).Exist()
	return !con
}
func Login(uname string, pwd string) (bool, bool) {
	var u User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	con := qs.Filter("Name", uname).Filter("Pwd", pwd)
	con.One(&u)
	return con.Exist(), u.IsAdmin
}
func AddUser(uname string, pwd string) (*User, error) {
	u := &User{Name:uname, Pwd:pwd, Pic:"/static/img/a.jpg"}
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err == nil {
		return u, err
	}
	return new(User), err
}

func CheckIsAdmin(uname string) bool {
	return GetUserInfo(uname).IsAdmin

}
func GetUserInfo(uname string) (User) {
	u := &User{Name:uname}
	o := orm.NewOrm()
	o.Read(u, "Name")
	return *u
}
func GetUserInfoByID(id string) (User) {
	uid, _ := strconv.Atoi(id)
	u := &User{Id:int64(uid)}
	o := orm.NewOrm()
	o.Read(u)
	return *u
}
func UpdateUserPic(id string, pic string) error {
	uid, _ := strconv.Atoi(id)
	u := &User{Id:int64(uid), Pic:pic}
	o := orm.NewOrm()
	_, err := o.Update(u, "Pic")
	return err

}
func UpdateUserPwd(id string, opwd string, newpwd string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("User").Filter("Id", id).Filter("Pwd", opwd)
	if qs.Exist() {
		_, err := qs.Update(orm.Params{
			"Pwd": newpwd,
		})
		if err == nil {
			return true
		}
	}

	return false
}

func UpdateById(id string, f common.UpdateUserForm) error {
	uid, _ := strconv.Atoi(id)
	u := &User{Id:int64(uid), NickName:f.NickName, Age:f.Age, Birth:f.Birth, Address:f.Address}
	if f.Sex == 1 {
		u.Sex = true
	} else {
		u.Sex = false
	}
	o := orm.NewOrm()
	_, err := o.Update(u, "NickName", "Age", "Sex", "Address", "Birth")
	return err

}

func GetUserList(page int) ([]User, int64, error) {
	var u []User
	o := orm.NewOrm()
	qs := o.QueryTable("User").OrderBy("Id")
	_, err := qs.Limit(15, 15 * page).All(&u)
	allnum, _ := qs.Count()

	return u, allnum, err

}