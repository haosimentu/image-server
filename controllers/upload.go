package controllers

import (
	"path"
	"image-server/utils"
	"os"
	"fmt"
	"github.com/disintegration/imaging"
	"strings"
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
	if ok {
		//创建目录
		uploadDir := os.Getenv("IMAGE_DIR")

		err := os.MkdirAll( uploadDir , 777)
		if err == nil {
			fullPath := uploadDir + fHeader.Filename

			//关闭上传的文件，不然的话会出现临时文件不能清除的情况
			defer mFile.Close()

			//保存原图
			err = controller.SaveToFile(fileKey, fullPath)
			if err == nil {
				//保存质量为75%的压缩图
				sImage, err := imaging.Open(fullPath)
				if err == nil {
					filenameWithSuffix := path.Base(fullPath)
					fileSuffix := path.Ext(filenameWithSuffix)
					filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
					fullPath = uploadDir + filenameOnly+"_q75"+fileSuffix

					err = utils.SaveImage(sImage, fullPath, fileSuffix, 75)

				}

				//返回响应
				data := map[string]string {"message": "success", "path": fullPath}
				controller.responseJson("0", data)
			}

			controller.warningResponse(fmt.Sprintf("文件保存失败. err=%v",err))
		}

		controller.warningResponse(fmt.Sprintf("%v", err))

	}

	controller.warningResponse(fmt.Sprintf("后缀名不符合上传要求, ext=%s", ext))

}
