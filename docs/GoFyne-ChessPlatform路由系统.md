# GoFyne-ChessPlatform 路由系统文档

## 概述
GoFyne-ChessPlatform 的路由系统在 Fyne 框架基础上提供了一种管理和切换不同用户界面视图的高效方法。本文档旨在说明如何在 GoFyne-ChessPlatform 项目中使用这个路由系统。

## 设计原则
- **模块化**: 确保视图和导航逻辑的解耦，增强代码模块化。
- **易于扩展**: 支持动态添加新的路由，便于应用未来的扩展和维护。
- **简洁接口**: 提供简单直观的接口，使得在应用中使用路由变得简单易行。

## 主要组件

### Router
- `Router` 结构体是路由系统的核心，负责管理路由映射和切换视图。
- 包含一个窗口实例和一个路由映射表。

### RouteConfig
- `RouteConfig` 结构体定义单个路由的配置，包括路由名称和对应的视图生成函数。

### 接口与方法
- `NewRouter(window fyne.Window) *Router`: 创建一个新的路由器实例。
- `RegisterRoute(name string, viewFunc func() fyne.CanvasObject)`: 注册一个路由。
- `GoTo(name string)`: 切换到指定名称的路由视图。
- `ConfigureRoutes(router *Router, routes []RouteConfig)`: 使用一组 `RouteConfig` 配置指定的路由器。

## 使用指南

### 初始化路由器
在应用启动时，使用 `NewRouter` 方法创建一个路由器实例，并关联到主窗口。

### 配置路由
使用 `RouteConfig` 结构体定义应用中的路由，并通过 `ConfigureRoutes` 方法配置到路由器。

### 视图切换
在应用的任何地方，通过调用路由器的 `GoTo` 方法来切换不同的视图。

## 示例代码

```go
// 创建路由器
router := route.NewRouter(window)

// 配置路由
	route.ConfigureRoutes(router, route.AppRoutes(router))

// 使用路由切换视图
router.GoTo("main")
```
### 注意事项
- 确保路由名称唯一，避免名称冲突。
- 在定义 `ViewFunc` 时，确保返回的 `fyne.CanvasObject` 正确代表了对应视图的布局和元素。

## 结论
GoFyne-ChessPlatform 的路由系统为 Fyne 应用提供了一种灵活且高效的界面管理方式。遵循本文档指南，你可以轻松地在应用中实现复杂的视图导航和管理。
