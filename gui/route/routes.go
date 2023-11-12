package route

import (
	"GFCP/gui/interfaces"
	"GFCP/gui/views"

	"fyne.io/fyne/v2"
)

// CreateRoute 创建并返回一个新的路由配置。
// 此函数接受一个路由器接口实例、路由的名称，以及一个创建视图的函数。
// 参数：
//
//	router   - 用于导航的路由器接口实例。
//	name     - 路由的唯一标识符。
//	viewFunc - 创建并返回 fyne.CanvasObject 的函数，用于生成路由的视图。
//
// 返回：
//
//	返回一个 RouteConfig 结构体，包含路由名称和视图生成函数。
func CreateRoute(router interfaces.RouterInterface, name string, viewFunc func(router interfaces.RouterInterface) fyne.CanvasObject) RouteConfig {
	return RouteConfig{
		Name: name,
		ViewFunc: func() fyne.CanvasObject {
			return viewFunc(router)
		},
	}
}

// AppRoutes 初始化并返回应用的路由配置列表。
// 此函数通过调用 CreateRoute 方法，为每个视图创建路由配置。
// 参数：
//
//	router - 用于导航的路由器接口实例。
//
// 返回：
//
//	返回 RouteConfig 类型的切片，包含应用中所有视图的路由配置。
func AppRoutes(router interfaces.RouterInterface) []RouteConfig {
	return []RouteConfig{
		CreateRoute(router, "main", views.NewMainView),
		CreateRoute(router, "settings", views.NewSettingsView),
		// ... 其他路由配置
	}
}
