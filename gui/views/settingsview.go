// gui/views/settingsview.go
package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsView(switchToMain func()) fyne.CanvasObject {
	backButton := widget.NewButton("Back to Main", switchToMain)
	label := widget.NewLabel("Settings")

	return container.NewVBox(
		label,
		backButton,
	)
}
