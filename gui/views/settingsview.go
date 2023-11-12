package views

import (
	"GFCP/gui/interfaces"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewSettingsView 创建设置视图
func NewSettingsView(router interfaces.RouterInterface) fyne.CanvasObject {
	backButton := widget.NewButton("Back to Main", func() {
		router.GoTo("main")
	})
	label := widget.NewLabel("Settings")

	return container.NewVBox(
		label,
		backButton,
	)
}
