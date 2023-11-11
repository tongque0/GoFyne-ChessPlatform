// gui/views/mainview.go
package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateMainView(switchToSettings func()) fyne.CanvasObject {
	settingsButton := widget.NewButton("Settings", switchToSettings)
	label := widget.NewLabel("Welcome to Chess Platform")

	return container.NewVBox(
		label,
		settingsButton,
	)
}
