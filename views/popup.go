package views

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewPopUp(message string, window fyne.Window) {
	popUpContent := widget.NewLabel(message)
	canvas := window.Canvas()
	popupWindow := widget.NewModalPopUp(popUpContent, canvas)

	go func() {
		<-time.After(time.Millisecond * 1500)
		popupWindow.Hide()
		canvas.Refresh(popupWindow)
	}()

	popupWindow.Show()
}
