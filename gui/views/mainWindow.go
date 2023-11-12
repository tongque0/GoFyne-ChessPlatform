package views

import (
	"GFCP/gui/interfaces"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainView(router interfaces.RouterInterface) fyne.CanvasObject {
	// 左侧边栏，这里仅用一个按钮表示
	leftSidebar := widget.NewButton("Settings", func() {
		router.GoTo("settings")
	})

	// 中间棋盘区域，这里用一个标签表示
	chessBoard := widget.NewLabel("Chess Board Here")

	// 右侧计时器区域，这里用一个标签表示
	rightTimer := widget.NewLabel("Timer Here")

	// 使用边界布局
	return container.NewBorder(nil, nil, leftSidebar, rightTimer, chessBoard)
}
