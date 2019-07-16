package controllers

import (
	"os"
	"image-server/utils"
	"fmt"
	"io/ioutil"
	"github.com/skip2/go-qrcode"
	"github.com/astaxie/beego"
)

type CodeController struct {
	Base
}

//二维码图
func (controller *CodeController) QCodeImage() {
	dir := os.Getenv("IMAGE_DIR")
	//filename_50*40.jpg
	url := controller.GetString("url")
	fileName := beego.Substr(utils.SHA256Encode(url), 10, 8)
	ext := ".png"

	fullPath := dir + fileName + ext
	exists := utils.CheckerFile(fullPath)
	if exists {
		file, err := ioutil.ReadFile(fullPath)
		if err != nil {
			controller.warningResponse(fmt.Sprintf("文件不存在， fileName=%s", fileName))
		}

		controller.Ctx.Output.Header("Content-Type", "image/jpeg")
		controller.Ctx.Output.Header("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"", fileName))
		controller.Ctx.ResponseWriter.Write(file)
	}

	err := qrcode.WriteFile(url, qrcode.Medium, 256, fullPath)
	if err != nil {
		controller.warningResponse("文件保存失败")
	}

	controller.Redirect("/show/"+fileName+ext, 302)
}

