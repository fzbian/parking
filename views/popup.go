package views

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"time"
)

func NewPopUp(message string, window fyne.Window) {
	popUpContent := widget.NewLabel(message)
	popupWindow := widget.ShowModalPopUp(popUpContent, window.Canvas())

	go func() {
		<-time.After(3 * time.Second)
		popupWindow.Hide()
	}()
}
