package delegate

import (
	http2 "github.com/wwbweibo/EasyRoute/src/http"
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetDefaultDelegate(contentRoot string) http2.RequestDelegate {
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
	defaultFileList := []string{
		"/index.html",
		"/index.htm",
		"/default.html",
		"/default.htm",
	}
	return func(ctx *context.Context) {
		for _, fileName := range defaultFileList {
			fileData, err := ioutil.ReadFile(wwwroot + fileName)
			if err != nil {
				continue
			}
			ctx.Response.WriteBody(fileData)
			ctx.Response.WriteHttpCode(http.StatusOK, "OK")
			if strings.Contains(ctx.Request.URL.Path, ".js") {
				ctx.Response.WriteHeader("Content-Type", []string{"application/x-javascript"})
			} else if strings.Contains(ctx.Request.URL.Path, ".svg") {
				ctx.Response.WriteHeader("Content-Type", []string{"image/svg+xml"})
			}
			return
		}
		NotFoundDelegate(ctx)
	}

}
