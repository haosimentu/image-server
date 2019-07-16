package controllers

import (
	"path"
	"image-server/utils"
	"os"
	"fmt"
)

type UploadController struct {
	Base
}

func (controller *UploadController) Upload() {
	//multipart.File, *multipart.FileHeader, error
	fileKey := "image"
	//获取上传的文件
	mFile, fHeader, _ := controller.GetFile(fileKey)
	ext := path.Ext(fHeader.Filename)

	//验证后缀名是否符合要求
	ok := utils.CheckExtValid(ext)
	if !ok {
		controller.warningResponse("后缀名不符合上传要求")
	}

	//创建目录
	uploadDir := os.Getenv("IMAGE_DIR")
	err := os.MkdirAll( uploadDir , 777)
	if err != nil {
		controller.warningResponse(fmt.Sprintf("%v", err))
	}

	fullPath := uploadDir + "/" + fHeader.Filename

	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer mFile.Close()

	err = controller.SaveToFile(fileKey, fullPath)
	if err != nil {
		controller.warningResponse(fmt.Sprintf("%v",err))
	}

	//返回响应
	data := map[string]string {"message": "success", "path": fullPath}
	controller.responseJson("0", data)
}
