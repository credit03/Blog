package utils

import (
	"strings"
	"fmt"
	"github.com/Unknwon/com"
	"os"
	"path"
	"github.com/sluu99/uuid"
	"time"
	"github.com/astaxie/beego"
)

func GetUserPicSavePath(filename string, uname string) string {

	arr := strings.Split(filename, ".")
	savefile := fmt.Sprint(uname, ".", arr[len(arr) - 1])
	savePath := fmt.Sprint("static/upload/pic/", savefile)
	if !com.IsExist(savePath) {
		os.MkdirAll(path.Dir(savePath), os.ModePerm)
	}
	return savePath;
}


func GetSavePath(filename string, types string) string {

	var id uuid.UUID = uuid.Rand()
	arr := strings.Split(filename, ".")
	savefile := fmt.Sprint(id.Hex(), ".", arr[len(arr) - 1])
	t := Substr(time.Now().Format(time.RFC3339), 0, 10)
	savePath := fmt.Sprint(types, t, "/", savefile)
	beego.Info("保存文件：", savePath)
	if !com.IsExist(savePath) {
		os.MkdirAll(path.Dir(savePath), os.ModePerm)
	}
	return savePath;
}


func GenerateType(t string, o string) string {
	var s string
	switch t {
	case "SectionType":
		switch o {
		case "-1":
			s = "全部文章"
		case "0":
			s = "生活版块"
		case "1":
			s = "科技版块"
		case "2":
			s = "其他版块"

		}


	case "ContentType":
		switch o {
		case "0":
			s = "文本文章"
		case "1":
			s = "图片文章"
		case "2":
			s = "视频文章"

		}
	}

	return s
}

func StringIsNotEmpty(val string) bool {
	if len(val) != 0 || val != "" {
		return true
	}
	return false
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func GetNo(no int) int {
	return no + 1
}