package routers

import (
	"image-server/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//页面路由
    beego.Router("/index.html", &controllers.MainController{},"get:Index")

	//接口路由
	beego.Router("/upload", &controllers.UploadController{},"post:Upload")

	//下载路由
	beego.Router("/download/:fileName", &controllers.DownloadController{}, "get:Download")

	//查看路由
	beego.Router("/show/:fileName", &controllers.ShowController{}, "get:Show")
	//beego.Router("/thumbnail/:fileName", &controllers.ShowController{}, "get:Thumbnail")
	//beego.Router("/blur/:fileName", &controllers.ShowController{}, "get:BlurImage")
	//beego.Router("/rotate/:fileName", &controllers.ShowController{}, "get:RotateImage")

	//二维码/水印等路由
	beego.Router("/qcode", &controllers.CodeController{}, "get:QCodeImage")

}
