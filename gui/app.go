package gui

import (
	"GFCP/gui/route"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartApp() {
	myApp := app.New()
	window := myApp.NewWindow("Chess Platform")
	window.Resize(fyne.NewSize(800, 600))
	router := route.NewRouter(window)

	// 使用 Router 类型的实例
	route.ConfigureRoutes(router, route.AppRoutes(router))

	router.GoTo("main")

	window.ShowAndRun()
}
