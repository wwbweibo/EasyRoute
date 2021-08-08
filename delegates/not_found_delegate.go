package delegates

import (
	http3 "github.com/wwbweibo/EasyRoute/http"
	"net/http"
)

var NotFoundDelegate = func(ctx *http3.Context) {
	_ = ctx.WritePlainText("404 Not Found", http.StatusNotFound)
}
