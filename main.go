package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/fzbian/parking/assets"
	"github.com/fzbian/parking/views"
)

func main() {

	// Create the app and the window.
	App := app.New()

	// Create the window.
	window := App.NewWindow("Parking José Max León Bilingual School")

	// Set the size of the window.
	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})

	// Set the size of the window to be fixed.
	window.SetFixedSize(true)
	// Center the window on the screen.
	window.CenterOnScreen()

	// Set the icon of the window.
	LogoResource := assets.GetIcon()
	window.SetIcon(LogoResource)

	// Get the main container and set it as the content of the window.
	MainContainer := views.GetMainContainer(window, LogoResource)

	// Show the window and run the app.
	window.SetContent(MainContainer)
	window.ShowAndRun()
}
