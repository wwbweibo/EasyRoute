package delegates

import (
	http2 "github.com/wwbweibo/EasyRoute/http"
)

type RequestDelegate func(ctx *http2.Context)
