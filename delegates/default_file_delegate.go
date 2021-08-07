package delegates

import (
	http3 "github.com/wwbweibo/EasyRoute/http"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetDefaultDelegate(contentRoot string) RequestDelegate {
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
	return func(ctx *http3.Context) {
		for _, fileName := range defaultFileList {
			fileData, err := ioutil.ReadFile(wwwroot + fileName)
			if err != nil {
				continue
			}

			ctx.Response.Write(fileData)
			ctx.Response.WriteHeader(http.StatusOK)
			if strings.Contains(ctx.Request.URL.Path, ".js") {
				ctx.Response.Header().Add("Content-Type", "application/x-javascript")
			} else if strings.Contains(ctx.Request.URL.Path, ".svg") {
				ctx.Response.Header().Set("Content-Type", "image/svg+xml")
			}
			return
		}
		NotFoundDelegate(ctx)
	}

}
