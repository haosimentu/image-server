package controllers

import (
	"os"
	"image-server/utils"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

type ShowController struct {
	Base
}

//展示图
func (controller *ShowController) Show() {
	//w=50&h=50&g=1&r=45&q=75&b=50
	var attrInfo utils.AttrInfo
	attrInfo.Width = controller.GetString("w")
	attrInfo.Height = controller.GetString("h")
	attrInfo.Gray = controller.GetString("g")
	attrInfo.Rotate = controller.GetString("r")
	attrInfo.Quality = controller.GetString("q")
	attrInfo.Blur = controller.GetString("b")

	//路径
	dir := os.Getenv("IMAGE_DIR")

	filenameWithSuffix :=	controller.Ctx.Input.Param(":fileName")
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	ok := utils.CheckExtValid(fileSuffix)

	if ok {
		filename := utils.GetFileName(filenameOnly, attrInfo)
		fullPath := dir + filename + fileSuffix
		exists := utils.CheckerFile(fullPath)
		if exists {
			file, err := ioutil.ReadFile(fullPath)
			if err != nil {
				controller.warningResponse(fmt.Sprintf("文件获取异常, err=%v", err))
			}

			controller.Ctx.Output.Header("Content-Type", "image/jpeg")
			controller.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", filename))
			controller.Ctx.ResponseWriter.Write(file)

		} else {
			//如果不存在，生成新图片并返回
			srcFullPath := dir + filenameWithSuffix
			srcExists := utils.CheckerFile(srcFullPath)
			if srcExists {
				//生成新文件
				srcImage, err := utils.LoadImage(srcFullPath)
				if err == nil {
					err = utils.SaveImage(srcImage, fullPath, fileSuffix, attrInfo)
					//返回响应
					if err == nil {
						file, _ := ioutil.ReadFile(fullPath)
						controller.Ctx.Output.Header("Content-Type", "image/jpeg")
						controller.Ctx.Output.Header("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"", filename))
						controller.Ctx.ResponseWriter.Write(file)
					} else {
						controller.warningResponse(fmt.Sprintf("文件生成异常, err=%v", err))
					}
				}
			} else {
				controller.warningResponse(fmt.Sprintf("原始文件不存在"))
			}
		}

	} else {
		controller.warningResponse(fmt.Sprintf("文件后缀不合法， fileSuffix=%s", fileSuffix))
	}

}


