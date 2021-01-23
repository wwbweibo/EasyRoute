package http

import "github.com/wwbweibo/EasyRoute/src/http/context"

type RequestDelegate func(ctx *context.Context)
