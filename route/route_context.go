package route

//
//import (
//	cctx "context"
//	"github.com/wwbweibo/EasyRoute/pkg/controllers"
//	"github.com/wwbweibo/EasyRoute/pkg/delegates"
//	iHttp "github.com/wwbweibo/EasyRoute/pkg/http"
//	TypeManagement2 "github.com/wwbweibo/EasyRoute/pkg/types"
//)
//
//var typeCollectionInstance = TypeManagement2.NewTypeCollect()
//
//type RouteContext struct {
//	controllers []*controllers.Controller // 添加到上下文中的控制器
//	// routeMap       map[string]routeMap         // 用于保存终结点和处理方法的映射
//	endPointTrie   *EndPointTrie                // 终结点前缀树
//	pipeline       Pipeline                     // 请求处理管道
//	app            delegates.RequestDelegate    // 最终的请求处理方法
//	typeCollection *TypeManagement2.TypeCollect // 内置类型字典
//	endpointPrefix string
//	server         interface{}
//	serverType     string
//	listenPort     string
//	listenAddress  string
//	ctx            cctx.Context
//}
//
//type paramMap struct {
//	paramName string
//	paramType string
//	source    string
//}
//
//func NewRouteContext(ctx cctx.Context) *RouteContext {
//	routeContext.ctx = ctx
//	return &routeContext
//}
//
//
//func (context *RouteContext) RegisterTypeByInstance(instance interface{}) {
//	context.typeCollection.Register(instance)
//}
//
//func (context *RouteContext) HandleRequest(ctx *iHttp.HttpContext) {
//	context.app(ctx)
//}
//
