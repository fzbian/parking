package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/fzbian/parking/assets"
	"github.com/fzbian/parking/views"
)

func main() {

	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")
	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	window.SetFixedSize(true)

	LogoResource := assets.GetIcon()
	window.SetIcon(LogoResource)

	MainContainer := views.GetMainContainer(window, LogoResource)

	window.SetContent(MainContainer)
	window.ShowAndRun()
}
