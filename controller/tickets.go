package controller

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"strings"
	"time"
)

var DevaVuFont = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"

// CreateEntryTicket creates input tickets validating the type of vehicle and assigning hours and values
func CreateEntryTicket(PlateNumber, VehicleType, SpotZone string, SpotID int) {

	// Create a new context with a white background
	ticket := gg.NewContext(200, 200)
	ticket.SetColor(color.White)
	ticket.Clear()

	// Add title
	TitleFontFace, err := gg.LoadFontFace(DevaVuFont, 20)
	ticket.SetFontFace(TitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Entry ticket", 100, 20, 0.5, 0.5)

	// Add subtitle
	SubtitleFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(SubtitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("José Max León Bilingual School", 100, 50, 0.5, 0.5)

	// Add vehicle information
	VehInfoFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(VehInfoFontFace)
	ticket.SetColor(color.Black)
	now := time.Now()
	horaActual := now.Format("15:04:05")
	vehicleInfo := fmt.Sprintf("Plate: %s\nType: %s\nSpot: %d\nZone: %s\nEntry time: %s", PlateNumber, VehicleType, SpotID, SpotZone, horaActual)
	DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)

	// Save image
	err = ticket.SavePNG("assets/tickets/entry/ " + PlateNumber + ".png")
	if err != nil {
		panic(err.Error())
	}

}

// CreateExitTicket creates output tickets validating the type of vehicle and assigning hours and values to be paid
func CreateExitTicket(PlateNumber, VehicleType, SpotZone, EntryTime string, SpotID int) {

	// Create a new context with a white background
	ticket := gg.NewContext(200, 200)
	ticket.SetColor(color.White)
	ticket.Clear()

	// Add title
	TitleFontFace, err := gg.LoadFontFace(DevaVuFont, 20)
	ticket.SetFontFace(TitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("Exit ticket", 100, 20, 0.5, 0.5)

	// Add subtitle
	SubtitleFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(SubtitleFontFace)
	ticket.SetColor(color.Black)
	ticket.DrawStringAnchored("José Max León Bilingual School", 100, 50, 0.5, 0.5)

	// Add vehicle information
	VehInfoFontFace, err := gg.LoadFontFace(DevaVuFont, 12)
	ticket.SetFontFace(VehInfoFontFace)
	ticket.SetColor(color.Black)
	now := time.Now()
	ExitTime := now.Format("15:04:05")

	// Calculate duration
	layout := "15:04:05"
	NewEntryTime, err1 := time.Parse(layout, EntryTime)
	NewExitTime, err2 := time.Parse(layout, ExitTime)

	if err1 != nil || err2 != nil {
		fmt.Println("Error when analyzing times: ", err1, err2)
		return
	}

	Duration := NewExitTime.Sub(NewEntryTime)

	// If is a provider vehicle, add amount payable
	if VehicleType == "PROVIDER" {
		if Duration.Minutes() > 30 {
			vehicleInfo := fmt.Sprintf("Plate: %s\nType: %s\nSpot: %d\nZone: %s\nEntry time: %s\nExit time: %s\nAmount payable: %s",
				PlateNumber,
				VehicleType,
				SpotID,
				SpotZone,
				EntryTime,
				ExitTime,
				"$5.000")
			DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)
		} else {
			vehicleInfo := fmt.Sprintf("Plate: %s\nType: %s\nSpot: %d\nZone: %s\nEntry time: %s\nExit time: %s",
				PlateNumber,
				VehicleType,
				SpotID,
				SpotZone,
				EntryTime,
				ExitTime)
			DrawTextWithNewlines(ticket, vehicleInfo, 10, 70, 180)
		}
	} else {
		vehicleInfo := fmt.Sprintf("Plate: %s\nType: %s\nSpot: %d\nZone: %s\nEntry time: %s\nExit time: %s",
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

// DrawTextWithNewlines draws line breaks in the given texts
func DrawTextWithNewlines(dc *gg.Context, text string, x, y, maxWidth float64) {

	// Load the DevaVuFont with a font size of 12 and handle any errors that occur
	font, err := gg.LoadFontFace(DevaVuFont, 12)
	if err != nil {
		panic(err)
	}

	// Set the gg.Context's font face to the font that was just loaded
	dc.SetFontFace(font)

	// Set the gg.Context's color to black
	dc.SetHexColor("#000000")

	// Split the text string by newline characters into an array of individual lines
	lines := strings.Split(text, "\n")

	// Loop through each line in the lines array
	for _, line := range lines {
		// Measure the width and height of the current line and store the height value in h
		_, h := dc.MeasureString(line)
		// Draw the current line with word wrapping, starting at the specified x and y coordinates, and using the specified maxWidth and line spacing of 1.5
		dc.DrawStringWrapped(line, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
		// Increment the y coordinate by the height of the current line multiplied by the line spacing of 1.5
		y += h * 1.5
	}

}
