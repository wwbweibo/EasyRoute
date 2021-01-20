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

func GetStaticFileMiddleware(withCache bool) route.Middleware {
	middleWare := func(next route.RequestDelegate) route.RequestDelegate {
		return func(ctx *context.Context) {
			// a static file should math *.* patten
			fmt.Println(ctx.Request.URL.Path)
			if strings.Contains(ctx.Request.URL.Path, ".") {
				dir, err := os.Getwd()
				if err != nil {
					ctx.Response.WriteBody([]byte("error to read file" + err.Error()))
					ctx.Response.WriteHttpCode(http.StatusInternalServerError)
					return
				}
				fileName := ctx.Request.URL.Path
				fileData, err := ioutil.ReadFile(dir + "/wwwroot" + fileName)
				if err != nil {
					ctx.Response.WriteBody([]byte("error to read file : " + dir + "/wwwroot" + fileName + "\r\n" + err.Error()))
					ctx.Response.WriteHttpCode(http.StatusInternalServerError)
					return
				}

				ctx.Response.WriteBody(fileData)
				ctx.Response.WriteHttpCode(http.StatusOK)
			}
		}
	}
	return middleWare
}
