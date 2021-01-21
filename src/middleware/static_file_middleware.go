package middleware

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetStaticFileMiddleware(contentRoot string, withCache bool) route.Middleware {
	var wwwroot string
	if contentRoot != "" {
		wwwroot = contentRoot
	} else {
		dir, err := os.Getwd()
		if err != nil {
			panic("error to read current work directory")
		}
		wwwroot = dir + "/wwwroot"
	}
	middleWare := func(next route.RequestDelegate) route.RequestDelegate {
		return func(ctx *context.Context) {
			// a static file should math *.* patten
			fmt.Println(ctx.Request.URL.Path)
			if strings.Contains(ctx.Request.URL.Path, ".") {
				fileName := ctx.Request.URL.Path
				fileData, err := ioutil.ReadFile(wwwroot + fileName)
				if err != nil {
					ctx.Response.WriteBody([]byte("error to read file : " + wwwroot + fileName + "\r\n" + err.Error()))
					ctx.Response.WriteHttpCode(http.StatusNotFound, "NotFound")
					return
				}

				ctx.Response.WriteBody(fileData)
				ctx.Response.WriteHttpCode(http.StatusOK, "OK")
			}
		}
	}
	return middleWare
}
