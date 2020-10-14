package main

import (
	"DataReceiver/Route"
	crypto2 "DataReceiver/crypto"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"reflect"
)

type ResultModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	routeContext := Route.NewRouteContext()

	controller := NewHomeController()
	controller.RegisterAsController(routeContext)
	routeContext.RouteParse()
	routeContext.Start(":8080")
}

type HomeController struct {
	Index func() string `Route:"/{Controller}/Index" method:"Get"`
}

func (self *HomeController) RegisterAsController(ctx *Route.RouteContext) {
	ctx.AddController(self)
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController() HomeController {
	return HomeController{
		Index: func() string {
			fmt.Println("enter index")
			return "Index"
		},
	}
}

func route(c *gin.Context) {
	path := c.Request.RequestURI
	if path == "/uploadData" {
		sn := c.PostForm("sn")
		bizType := c.PostForm("bizType")
		content := c.PostForm("content")
		file, _ := c.FormFile("file")
		ext := c.PostForm("ext")
		sign := c.PostForm("sign")
		c.JSON(http.StatusOK, uploadData(sn, bizType, content, file, ext, sign))
	} else {
		c.String(http.StatusNotFound, "404 NotFind")
	}
}

func uploadData(sn, bizType, content string, file *multipart.FileHeader, ext string, sign string) ResultModel {
	result := ResultModel{}
	result.Code = 500
	result.Message = "Error"

	if sn == "" {
		result.Message = "Sn不能为空"
		return result
	}
	if bizType == "" {
		result.Message = "业务类型不能为空"
		return result
	}
	if content == "" {
		result.Message = "Content不能为空"
		return result
	}
	checkData := sn + bizType + content
	var bytes = []byte(checkData)
	if file != nil {
		fileContent, err := file.Open()
		if err != nil {
			result.Message = "文件读取错误"
			return result
		}
		fileBytes := make([]byte, file.Size)
		fileContent.Read(fileBytes)
		bytes = append(fileBytes, bytes...)
	}
	bytes = append(bytes, []byte(ext)...)
	data := crypto2.EncryptData(bytes)
	computedSign := data.HMacMd5("123")
	if sign == "" || sign != computedSign {
		result.Code = 401
		result.Message = "签名校验失败"
		return result
	}
	result.Code = 200
	result.Message = "Success"
	return result
}
