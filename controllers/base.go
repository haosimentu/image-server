package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type Base struct {
	beego.Controller
}

func (base *Base) ValidatorAuto(frontendData interface{}) {
	if err := json.Unmarshal(base.Ctx.Input.RequestBody, &frontendData); err != nil {
		base.errorHandler(err)
	}
	defaultMessage := map[string]string{
		"Required":     "不能为空",
		"Min":          "不能小于%d",
		"Max":          "不能大于%d",
		"Range":        "取值必须在%d到%d之间",
		"MinSize":      "长度不能小于%d",
		"MaxSize":      "长度不能大于%d",
		"Length":       "长度必须等于%d",
		"Alpha":        "必须是字母",
		"Numeric":      "必须是数字",
		"AlphaNumeric": "必须是字母或者数字",
		"Match":        "必须出现 %s 关键字",
		"NoMatch":      "不能出现 %s 关键字",
		"AlphaDash":    "必须是字母，数组或者横线(-)",
		"Email":        "不合法的邮箱地址",
		"IP":           "不合法的IP",
		"Base64":       "不合法的Base64编码格式",
		"Mobile":       "不合法的手机号",
		"Tel":          "不合法的电话号码",
		"Phone":        "不合法的手机号",
		"ZipCode":      "不合法的邮编",
	}
	validation.SetDefaultMessage(defaultMessage)
	validate := validation.Validation{}

	isValid, err := validate.Valid(frontendData)
	if err != nil {
		base.errorHandler(err)
	}

	if !isValid {
		base.warningResponse(validate.Errors[0].Message)
	}
}

func (base *Base) errorHandler(err error) {
	logs.Info(err)
	data := make(map[string]string)
	data["message"] = "系统出错"
	base.responseJson("9", data)
}

func (base *Base) successResponse() {
	data := make(map[string]string)
	data["message"] = ""
	base.responseJson("0", data)
}

func (base *Base) warningResponse(message string) {
	data := make(map[string]string)
	data["message"] = message
	base.responseJson("5", data)
}

func (base *Base) responseJson(code string, data map[string]string) {
	data["code"] = code
	base.Data["json"] = data
	base.ServeJSON()
	base.StopRun()
}

func (base *Base) Prepare() {
	base.EnableXSRF = false
}
