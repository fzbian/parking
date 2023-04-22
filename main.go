package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/fzbian/parking/assets"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/views"
)

func main() {

	//zone, err := controller.GetTotalTimeByZone(models.Spot{Zone: "A"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(zone)

	zone, err := controller.GetMostUsedZone()
	if err != nil {
		panic(err)
	}
	fmt.Print(zone)

	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")
	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	window.SetFixedSize(true)
	window.CenterOnScreen()

	LogoResource := assets.GetIcon()
	window.SetIcon(LogoResource)

	MainContainer := views.GetMainContainer(window, LogoResource)

	window.SetContent(MainContainer)
	window.ShowAndRun()
}

/* TODO
- Mostrar por sonido cuando el parqueadero este por llenarse y cuando este lleno
*/
