package controllers

import (
	"os"
	"image-server/utils"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/astaxie/beego"
	"image"
)

type ShowController struct {
	Base
}

//展示图
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

//模糊图
func (controller *ShowController) BlurImage() {
	dir := os.Getenv("IMAGE_DIR")
	//filename_50*40.jpg
	fileName :=	controller.Ctx.Input.Param(":fileName")
	ext := filepath.Ext(fileName)
	ok := utils.CheckExtValid(ext)

	if ok {

		fullPath := dir + fileName
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

		sIdx := strings.LastIndex(fileName, "_")
		cIdx := strings.LastIndex(fileName, ".")

		sb := beego.Substr(fileName, 0, sIdx)
		ss := beego.Substr(fileName, sIdx+1, cIdx-(sIdx+1))

		if len(ss) <= 0 {
			controller.warningResponse("图片加载失败")
		}

		tempPath := dir + sb + ext
		srcImage, err := utils.LoadImage(tempPath)

		if err != nil {
			controller.warningResponse("图片加载失败")
		}

		dev , err := strconv.ParseFloat(ss, 2)
		if err != nil {
			dev = 1.0
		}

		bound := srcImage.Bounds()
		dx := bound.Dx()
		dy := bound.Dy()
		destImage := image.NewRGBA(image.Rect(0, 0, dx, dy))
		err = graphics.Blur(destImage, srcImage, &graphics.BlurOptions{StdDev: dev})
		err = utils.SaveImage(fullPath, destImage, ext)
		if err != nil {
			controller.warningResponse(fmt.Sprintf("文件保存错误， err=%", err))
		}


		controller.Redirect("/show/"+fileName, 302)
	}

	controller.warningResponse(fmt.Sprintf("文件不合法， ext=%s", ext))
}

//缩略图
func (controller *ShowController) Thumbnail() {
	dir := os.Getenv("IMAGE_DIR")
	//filename_50*40.jpg
	fileName :=	controller.Ctx.Input.Param(":fileName")
	ext := filepath.Ext(fileName)
	ok := utils.CheckExtValid(ext)

	if ok {

		fullPath := dir + fileName
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

		sIdx := strings.LastIndex(fileName, "_")
		cIdx := strings.LastIndex(fileName, ".")

		sb := beego.Substr(fileName, 0, sIdx)
		ss := beego.Substr(fileName, sIdx+1, cIdx-(sIdx+1))

		if len(ss) <= 0 {
			controller.warningResponse("图片加载失败")
		}

		tempPath := dir + sb + ext
		srcImage, err := utils.LoadImage(tempPath)
		if err != nil {
			controller.warningResponse("图片加载失败")
		}

		bound := srcImage.Bounds()
		dx := bound.Dx()
		dy := bound.Dy()

		var newX, newY int
		//ss = strings.Replace(ss, "*", "x", 1)
		xy := strings.Split(ss, "x")
		if len(xy) > 1 {
			newX, _ = strconv.Atoi(xy[0])
			newY, _ = strconv.Atoi(xy[1])
		} else {
			newX = dx
			newY = dy
		}
		// 缩略图的大小
		destImage := image.NewRGBA(image.Rect(0, 0, newX, newY))
		err = graphics.Scale(destImage, srcImage)
		err = utils.SaveImage(fullPath, destImage, ext)
		if err != nil {
			controller.warningResponse(fmt.Sprintf("文件保存错误， err=%", err))
		}

		//controller.Ctx.Output.Header("Content-Type", "image/jpeg")
		//controller.Ctx.Output.Header("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"", fileName))

		controller.Redirect("/show/"+fileName, 302)
	}

	controller.warningResponse(fmt.Sprintf("文件不合法， ext=%s", ext))
}

//旋转图
func (controller *ShowController) RotateImage() {
	dir := os.Getenv("IMAGE_DIR")
	//filename_50*40.jpg
	fileName :=	controller.Ctx.Input.Param(":fileName")
	ext := filepath.Ext(fileName)
	ok := utils.CheckExtValid(ext)

	if ok {

		fullPath := dir + fileName
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

		sIdx := strings.LastIndex(fileName, "_")
		cIdx := strings.LastIndex(fileName, ".")

		sb := beego.Substr(fileName, 0, sIdx)
		ss := beego.Substr(fileName, sIdx+1, cIdx-(sIdx+1))

		if len(ss) <= 0 {
			controller.warningResponse("图片加载失败")
		}

		tempPath := dir + sb + ext
		srcImage, err := utils.LoadImage(tempPath)

		if err != nil {
			controller.warningResponse("图片加载失败")
		}

		op, err := strconv.ParseFloat(ss, 3)
		if err != nil {
			op = 0.0
		}

		bound := srcImage.Bounds()
		dx := bound.Dx()
		dy := bound.Dy()
		destImage := image.NewRGBA(image.Rect(0, 0, dx, dy))
		err = graphics.Rotate(destImage, srcImage, &graphics.RotateOptions{op})
		err = utils.SaveImage(fullPath, destImage, ext)
		if err != nil {
			controller.warningResponse(fmt.Sprintf("文件保存错误， err=%", err))
		}


		controller.Redirect("/show/"+fileName, 302)
	}

	controller.warningResponse(fmt.Sprintf("文件不合法， ext=%s", ext))
}


