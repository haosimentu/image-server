package controllers

import (
	"os"
	"image-server/utils"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

type ShowController struct {
	Base
}

//展示原图
func (controller *ShowController) Show() {
	dir := os.Getenv("IMAGE_DIR")
	fileName :=	controller.Ctx.Input.Param(":fileName")
	ext := filepath.Ext(fileName)
	ok := utils.CheckExtValid(ext)

	if ok {
		fullPath := dir + fileName
		exists := utils.CheckerFile(fullPath)
		if exists {
			file, err := ioutil.ReadFile(fullPath)
			if err != nil {
				controller.warningResponse(fmt.Sprintf("文件获取异常, err=%v", err))
			}

			controller.Ctx.Output.Header("Content-Type", "image/jpeg")
			controller.Ctx.Output.Header("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"", fileName))
			controller.Ctx.ResponseWriter.Write(file)
		}

		controller.warningResponse("文件不存在")
	}

	controller.warningResponse(fmt.Sprintf("文件后缀不合法， ext=%s", ext))
}


