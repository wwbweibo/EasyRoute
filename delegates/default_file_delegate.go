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
			contentType := "text/html"
			if strings.Contains(ctx.Request.URL.Path, ".js") {
				contentType = "application/x-javascript"
			} else if strings.Contains(ctx.Request.URL.Path, ".svg") {
				contentType = "image/svg+xml"
			}
			_ = ctx.Write(fileData, http.StatusOK, contentType)
			return
		}
		NotFoundDelegate(ctx)
	}

}
