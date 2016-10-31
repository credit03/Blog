package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"fmt"
	"../utils"
)

type UploadController struct {
	beego.Controller
}

const (
	IamgePath = "static/upload/iamges/"
	VideoPath = "static/upload/videos/"
)

func (this *UploadController) Get() {
	beego.Info("文件上传get")
	this.Ctx.WriteString("/static/img/1.jpg")
	return
}

func (this *UploadController) ImageUpload() {
	f, h, err := this.GetFile("upload")
	defer f.Close()
	caback := this.Input().Get("CKEditorFuncNum")
	if err == nil {
		//2MB
		buf := make([]byte, 2 * 1024 * 1024)
		var l int
		l, err = f.Read(buf)
		if err == nil {
			//限制上传文件大小
			if l < (2 * 1024 * 1024) {
				beego.Info("文件大小", l, len(buf))
				p := utils.GetSavePath(h.Filename, IamgePath)
				err = this.SaveToFile("upload", p)
				if err == nil {
					beego.Info("插入数据库")
					if models.ImageTempFileAdd(p) == nil {
						out := fmt.Sprint("<script type=\"text/JavaScript\">", "window.parent.CKEDITOR.tools.callFunction(", caback, ",'/", p, "','')</script>")
						this.Ctx.WriteString(out);
						return
					}
				}
			} else {
				out := fmt.Sprint("<script type=\"text/JavaScript\">alert('上传文件过大，请上传小于2MB的图片...')</script>")
				this.Ctx.WriteString(out);
				return
			}

		}
	}
	out := fmt.Sprint("<script type=\"text/JavaScript\">alert('上传失败...')</script>")
	this.Ctx.WriteString(out);


	return
}

func (this *UploadController) Prepare() {
	CheckUserLogin(&(this.Controller))
	//XSRF 之前是全局设置的一个参数,如果设置了那么所有的API请求都会进行验证,但是有些时候API 逻辑是不需要进行验证的,因此现在支持在controller 级别设置屏蔽:
	this.EnableXSRF = false
}


/**
   备用接口，由于视频是流媒体
 */
func (this *UploadController) VideoUpload() {
	f, h, err := this.GetFile("videoupload")
	defer f.Close()
	caback := this.Input().Get("CKEditorFuncNum")
	if err == nil {

		p := utils.GetSavePath(h.Filename, VideoPath)
		err = this.SaveToFile("videoupload", p)
		if err == nil {
			if models.VideoTempFileAdd(p) == nil {
				out := fmt.Sprint("<script type=\"text/JavaScript\">", "window.parent.CKEDITOR.tools.callFunction(", caback, ",'/", p, "','')</script>")
				this.Ctx.WriteString(out);
				return
			}

		}
	}

	out := fmt.Sprint("<script type=\"text/JavaScript\">alert('上传失败...')</script>")
	this.Ctx.WriteString(out);

}
