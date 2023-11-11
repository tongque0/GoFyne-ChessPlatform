// gui/app.go
package gui

import (
	"GFCP/gui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// SetupUI 构建并返回应用程序的主界面
func SetupUI(app fyne.App) fyne.Window {
	window := app.NewWindow("Chess Platform")

	var showMainView, showSettingsView func()

	showMainView = func() {
		mainView := views.CreateMainView(showSettingsView)
		window.SetContent(mainView)
	}

	showSettingsView = func() {
		settingsView := views.CreateSettingsView(showMainView)
		window.SetContent(settingsView)
	}

	showMainView() // 默认显示主界面

	window.Resize(fyne.NewSize(600, 400))
	return window
}
func StartApp() {
	myApp := app.New()
	mainWindow := SetupUI(myApp)

	mainWindow.ShowAndRun()
}
