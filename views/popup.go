package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

func NewPopUp(message string, window fyne.Window) {
	popUpContent := widget.NewLabel(message)
	canvas := window.Canvas()
	popupWindow := widget.NewModalPopUp(popUpContent, canvas)

	go func() {
		<-time.After(3 * time.Second)
		popupWindow.Hide()
		canvas.Refresh(popupWindow)
	}()

	popupWindow.Show()
}
