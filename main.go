package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/fzbian/parking/assets"
	"github.com/fzbian/parking/views"
)

func main() {

	testdata := views.GetData()
	jsonData, err := json.Marshal(testdata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	app := app.New()

	window := app.NewWindow("Parqueadero Colegio Jose Max Leon")
	window.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	window.SetFixedSize(true)

	LogoResource := assets.GetIcon()
	window.SetIcon(LogoResource)

	LogoObject := canvas.NewImageFromResource(LogoResource)
	LogoContainer := container.New(layout.NewGridWrapLayout(fyne.Size{
		Width:  175,
		Height: 175,
	}), LogoObject)

	VehicleTable := views.GetVehiclesTable()
	AddVehicleButton, ExitVehicleButton := views.LeftButtons(window)
	ExportRecordsButton, ExitWindowButton := views.RightButtons(window)

	LeftContainer := container.New(layout.NewVBoxLayout(),
		AddVehicleButton,
		ExitVehicleButton,
		LogoContainer,
	)

	MidContainer := container.New(layout.NewGridWrapLayout(
		fyne.Size{
			Width:  900,
			Height: 650,
		}), VehicleTable)

	RightContainer := container.New(layout.NewVBoxLayout(),
		ExportRecordsButton,
		ExitWindowButton,
	)

	MainContainer := container.New(layout.NewHBoxLayout(),
		LeftContainer,
		MidContainer,
		RightContainer,
	)

	window.SetContent(MainContainer)
	window.ShowAndRun()
}
