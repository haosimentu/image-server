package routers

import (
	"image-server/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{},"get:Index")
	beego.Router("/upload", &controllers.UploadController{},"post:Upload")
	beego.Router("/download/:fileName", &controllers.DownloadController{}, "get:Download")
	beego.Router("/show/:fileName", &controllers.ShowController{}, "get:Show")
	beego.Router("/thumbnail/:fileName", &controllers.ShowController{}, "get:Thumbnail")
	beego.Router("/blur/:fileName", &controllers.ShowController{}, "get:BlurImage")
	beego.Router("/rotate/:fileName", &controllers.ShowController{}, "get:RotateImage")
	beego.Router("/qcode", &controllers.CodeController{}, "get:QCodeImage")


	//注解
	//beego.Include(&controllers.MainController{})
	//beego.Include(&controllers.UploadController{})

}
