package route

import (
	"GFCP/gui/interfaces"

	"fyne.io/fyne/v2"
)

// Router 是一个用于管理应用视图和路由的结构体。
// 它包含一个映射，用于存储路由名称和对应的视图生成函数，
// 以及一个 fyne.Window 实例，用于显示当前活动的视图。
type Router struct {
	routes map[string]func() fyne.CanvasObject
	window fyne.Window
}

// 确保 Router 结构体实现了 RouterInterface
var _ interfaces.RouterInterface = &Router{}

// NewRouter 创建一个新的 Router 实例。
// 参数：
//
//	window - fyne.Window 实例，用于展示视图。
//
// 返回：
//
//	返回一个初始化的 Router 实例。
func NewRouter(window fyne.Window) *Router {
	return &Router{
		routes: make(map[string]func() fyne.CanvasObject),
		window: window,
	}
}

// RegisterRoute 注册一个新的路由到路由器。
// 参数：
//
//	name     - 路由的唯一标识符。
//	viewFunc - 用于生成视图的函数。
func (r *Router) RegisterRoute(name string, viewFunc func() fyne.CanvasObject) {
	r.routes[name] = viewFunc
}

// GoTo 根据给定的路由名称切换视图。
// 如果给定的路由名称存在，则切换到对应的视图。
// 参数：
//
//	name - 路由名称。
func (r *Router) GoTo(name string) {
	if viewFunc, exists := r.routes[name]; exists {
		r.window.SetContent(viewFunc())
	}
}

// RouteConfig 定义单个路由的配置结构。
// 它包含路由的名称和一个生成视图的函数。
type RouteConfig struct {
	Name     string
	ViewFunc func() fyne.CanvasObject
}

// ConfigureRoutes 使用一组 RouteConfig 来配置路由器。
// 此函数将每个 RouteConfig 中的路由添加到路由器中。
// 参数：
//
//	router - 路由器实例。
//	routes - 一组路由配置。
func ConfigureRoutes(router *Router, routes []RouteConfig) {
	for _, route := range routes {
		router.RegisterRoute(route.Name, route.ViewFunc)
	}
}
