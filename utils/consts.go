package utils

import (
	"image"
	"os"
	"crypto"
	"encoding/hex"
	"github.com/disintegration/imaging"
	"image/png"
	"strings"
	"strconv"
	"image/color"
)

type AttrInfo struct {
	Width, Height, Gray, Rotate, Blur, Quality string
}

const (
	SEP string = "_"
	WIDTH string = "w"
	WIDTHSEP string = "_w"
	HEIGH string = "h"
	HEIGHSEP string = "_h"
	GRAY string = "g"
	GRAYSEP string = "_g"
	ROTATE string = "r"
	ROTATESEP string = "_r"
	QUALITY string = "q"
	QUALITYSEP string = "_q"
	BLUR string = "b"
	BLURSEP string = "_b"
)

var AllowExtMap map[string]bool = map[string]bool {
	".jpg":true,
	".jpeg":true,
	".png":true,
}

func CheckExtValid(ext string) (bool) {
	_,ok := AllowExtMap[ext]
	return ok
}


func LoadImage(filePath string) (img image.Image, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	img, _, err = image.Decode(file)

	return
}

func CheckerFile(fPath string) (bool) {
	_, err := os.Stat(fPath)
	if err == nil {
		return true
	} else {
		return false
	}
}

func SaveImage(img image.Image, path string, ext string, attrInfo AttrInfo) (err error) {

	//生成缩略图，宽x,高0表示等比例放缩
	var width, height int
	if attrInfo.Width == "" {
		if attrInfo.Height != "" {
			height, _ = strconv.Atoi(attrInfo.Height)
			img = imaging.Resize(img, width, height, imaging.Lanczos)
		}
	} else {

		width, _ = strconv.Atoi(attrInfo.Width)
		if attrInfo.Height != "" {
			height, _ = strconv.Atoi(attrInfo.Height)
		}

		img = imaging.Resize(img, width, height, imaging.Lanczos)
	}

	//旋转
	if attrInfo.Rotate != "" {
		rotate , _ := strconv.ParseFloat(attrInfo.Rotate, 2)
		if rotate > 0 {
			img = imaging.Rotate(img, rotate, color.Opaque)
		}
	}

	//模糊
	if attrInfo.Blur != "" {
		blur , _ := strconv.ParseFloat(attrInfo.Blur, 2)
		if blur > 0 {
			img = imaging.Blur(img, blur)
		}
	}

	//灰度
	if attrInfo.Gray != "" {
		gray , _ := strconv.ParseFloat(attrInfo.Gray, 2)
		if gray > 0 {
			img = imaging.AdjustSaturation(img, gray)
		}
	}

	//压缩比
	quality := 100
	if attrInfo.Quality != "" {
		quality, _ = strconv.Atoi(attrInfo.Quality)
	}

	switch ext {
	case ".png":
		err = imaging.Save(img, path,imaging.PNGCompressionLevel(png.CompressionLevel(quality)))
		break
	case ".jpeg":
		err = imaging.Save(img, path,imaging.JPEGQuality(quality))
		break
	case ".jpg":
		err = imaging.Save(img, path,imaging.JPEGQuality(quality))
		break
	}

	return
}

func SHA256Encode(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}

//aa$h500$w500$g1$r45$q75
func FromFileName(filename string) (AttrInfo) {
	attrs := strings.Split(filename, SEP)

	var attrInfo AttrInfo

	for _, attr := range attrs {
		switch attr[0:1] {
		case WIDTH :
			attrInfo.Width = attr[1:]
			break
		case HEIGH:
			attrInfo.Height = attr[1:]
			break
		case GRAY:
			attrInfo.Gray = attr[1:]
			break
		case ROTATE:
			attrInfo.Rotate = attr[1:]
			break
		case QUALITY:
			attrInfo.Quality = attr[1:]
			break
		case BLUR:
			attrInfo.Blur = attr[1:]
			break
		default:
			break
		
		}
	}

	return attrInfo
}

//w=500&h=500&g=1&r=45&q=75&f=jpeg
//aa$h500$w500$g1$r45$q75
func GetFileName(filenameOnly string, info AttrInfo) (string) {
	if info.Width != "" {
		filenameOnly = filenameOnly + WIDTHSEP + info.Width
	}

	if info.Height != "" {
		filenameOnly = filenameOnly + HEIGHSEP + info.Height
	}

	if info.Gray != "" {
		filenameOnly = filenameOnly + GRAYSEP + info.Gray
	}

	if info.Rotate != "" {
		filenameOnly = filenameOnly + ROTATESEP + info.Rotate
	}

	if info.Quality != "" {
		filenameOnly = filenameOnly + QUALITYSEP + info.Quality
	}

	if info.Blur != "" {
		filenameOnly = filenameOnly + BLURSEP + info.Blur
	}

	return filenameOnly
}