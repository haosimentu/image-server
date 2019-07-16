package main

import (
	_ "image-server/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
	"log"
	"image-server/utils"
)

func main() {
	var pwd = filepath.Dir(os.Args[0])
	var envFile = pwd + string(filepath.Separator) + ".env"
	println(os.Getenv("UPLOAD_DIR"))
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("can't find .env file")
		return
	}

	//跨域配置
	utils.CorsHandler()

	logs.SetLogger("console")

	// 运行
	beego.Run(":" + os.Getenv("APP_PORT"))

}
