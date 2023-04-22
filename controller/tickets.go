package controller

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"strings"
	"time"
)

var DevaVuFont = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

func CreateEntryTicket(PlateNumber, VehicleType, SpotZone string, SpotID int) {

	// Create a new context with a white background
	ticket := gg.NewContext(200, 200)
	ticket.SetColor(color.White)
	ticket.Clear()

	// Add title
	TitleFontFace, err := gg.LoadFontFace(DevaVuFont, 20)
	ticket.SetFontFace(TitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Ticket de entrada", 100, 20, 0.5, 0.5)

	// Add subtitle
	SubtitleFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(SubtitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Colegio Bilingüe José Max León", 100, 50, 0.5, 0.5)

	// Add vehicle information
	VehInfoFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(VehInfoFontFace)
	ticket.SetColor(color.Black)
	now := time.Now()
	horaActual := now.Format("15:04:05")
	vehicleInfo := fmt.Sprintf("Placa: %s\nTipo: %s\nBahia: %d\nZona: %s\nHora de entrada: %s", PlateNumber, VehicleType, SpotID, SpotZone, horaActual)
	DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)

	// Save image
	err = ticket.SavePNG("assets/tickets/entry/ " + PlateNumber + ".png")
	if err != nil {
		panic(err.Error())
	}
}

func CreateExitTicket(PlateNumber, VehicleType, SpotZone, EntryTime string, SpotID int) {

	// Create a new context with a white background
	ticket := gg.NewContext(200, 200)
	ticket.SetColor(color.White)
	ticket.Clear()

	// Add title
	TitleFontFace, err := gg.LoadFontFace(DevaVuFont, 20)
	ticket.SetFontFace(TitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Ticket de salida", 100, 20, 0.5, 0.5)

	// Add subtitle
	SubtitleFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(SubtitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Colegio Bilingüe José Max León", 100, 50, 0.5, 0.5)

	// Add vehicle information
	VehInfoFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(VehInfoFontFace)
	ticket.SetColor(color.Black)
	now := time.Now()
	ExitTime := now.Format("15:04:05")

	layout := "15:04:05"
	NewEntryTime, err1 := time.Parse(layout, EntryTime)
	NewExitTime, err2 := time.Parse(layout, ExitTime)

	if err1 != nil || err2 != nil {
		fmt.Println("Error al analizar los tiempos:", err1, err2)
		return
	}

	Duration := NewExitTime.Sub(NewEntryTime)

	if VehicleType == "PROVEEDOR" {
		if Duration.Minutes() > 30 {
			vehicleInfo := fmt.Sprintf("Placa: %s\nTipo: %s\nBahia: %d\nZona: %s\nHora de entrada: %s\nHora de salida: %s\nValor a pagar: %s",
				PlateNumber,
				VehicleType,
				SpotID,
				SpotZone,
				EntryTime,
				ExitTime,
				"$5.000")
			DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)
		} else {
			vehicleInfo := fmt.Sprintf("Placa: %s\nTipo: %s\nBahia: %d\nZona: %s\nHora de entrada: %s\nHora de salida: %s",
				PlateNumber,
				VehicleType,
				SpotID,
				SpotZone,
				EntryTime,
				ExitTime)
			DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)
		}
	} else {
		vehicleInfo := fmt.Sprintf("Placa: %s\nTipo: %s\nBahia: %d\nZona: %s\nHora de entrada: %s\nHora de salida: %s",
			PlateNumber,
			VehicleType,
			SpotID,
			SpotZone,
			EntryTime,
			ExitTime)
		DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)
	}

	// Save image
	err = ticket.SavePNG("assets/tickets/exit/ " + PlateNumber + ".png")
	if err != nil {
		panic(err.Error())
	}
}

func DrawTextWithNewlines(dc *gg.Context, text string, x, y, maxWidth float64) {
	font, err := gg.LoadFontFace(DevaVuFont, 12)
	if err != nil {
		panic(err)
	}
	dc.SetFontFace(font)
	dc.SetHexColor("#000000")

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		_, h := dc.MeasureString(line)
		dc.DrawStringWrapped(line, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
		y += h * 1.5
	}
}
