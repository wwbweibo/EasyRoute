package delegates

type Middleware func(next RequestDelegate) RequestDelegate
