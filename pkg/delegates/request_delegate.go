package delegates

import (
	"github.com/wwbweibo/EasyRoute/pkg/http"
)

type RequestDelegate func(ctx *http.HttpContext)
