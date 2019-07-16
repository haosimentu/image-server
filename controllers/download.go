package controllers

import (
	"os"
	"image-server/utils"
	"path/filepath"
	"fmt"
)

type DownloadController struct {
	Base
}

func (controller *DownloadController) Download() {
	dir := os.Getenv("IMAGE_DIR")
	fileName :=	controller.Ctx.Input.Param(":fileName")
	ext := filepath.Ext(fileName)
	ok := utils.CheckExtValid(ext)
	if ok {
		fPath := dir + fileName
		exists := utils.CheckerFile(fPath)
		if exists {
			controller.Ctx.Output.Download(fPath)
		}
	}

	controller.warningResponse(fmt.Sprintf("文件不存在或后缀不合法， ext=%s", ext))
}
