package utils

import (
	"image"
	"os"
	"crypto"
	"encoding/hex"
	"github.com/disintegration/imaging"
	"image/png"
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

func SaveImage(img image.Image, path string, ext string, quality int) (err error) {
	//imgfile, err := os.Create(path)
	//defer imgfile.Close()

	//if err != nil {
	//	return
	//}

	switch ext {
	case ".png":
		//err = png.Encode(imgfile, img)
		err = imaging.Save(img, path, imaging.PNGCompressionLevel(png.CompressionLevel(quality)))
		break
	case ".jpeg":
		err = imaging.Save(img, path, imaging.JPEGQuality(quality))
		break
	case ".jpg":
		//err = jpeg.Encode(imgfile, img, &jpeg.Options{100})
		err = imaging.Save(img, path, imaging.JPEGQuality(quality))
		break
	}

	return
}

func SHA256Encode(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}