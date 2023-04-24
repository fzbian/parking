package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
)

// GetTable returns a *widget.Table with the data of the spots.
func GetTable() *widget.Table {

	// Get the spots from the database and store them in a slice.
	var spots []models.Spot
	utils.Db.Find(&spots)

	// Create a slice of slices to store the data of the spots.
	var data [][]string

	// Iterate over the spots and append the data to the data slice.
	for _, u := range spots {
		// If the spot is in use, get the plate number of the vehicle that is parked in the spot.
		if u.InUse {
			plateNumber := controller.GetVehiclePlateNumberBySpotId(u.ID)
			data = append(data, []string{u.GetIDString(), u.Type, u.Zone, "Occupied", plateNumber})
			// If the spot is free, append an empty string to the plate number column.
		} else {
			data = append(data, []string{u.GetIDString(), u.Type, u.Zone, "Free", ""})
		}
	}

	// Create the table with the data.
	table := widget.NewTable(
		// The function that returns the number of rows and columns.
		func() (int, int) {
			return len(data) + 1, len(data[0])
		},
		// The function that returns the widget that will be displayed in the cell.
		func() fyne.CanvasObject {
			return widget.NewLabel("TableCell")
		},
		// The function that sets the data of the cell.
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				switch i.Col {
				case 0:
					o.(*widget.Label).SetText("ID")
				case 1:
					o.(*widget.Label).SetText("Type")
				case 2:
					o.(*widget.Label).SetText("Zone")
				case 3:
					o.(*widget.Label).SetText("Status")
				case 4:
					o.(*widget.Label).SetText("Plate Number")
				}
				// If the row is not 0, set the text of the cell to the data of the spot.
			} else {
				o.(*widget.Label).SetText(data[i.Row-1][i.Col])
			}
		},
	)

	// Adjust the width of the columns.
	table.SetColumnWidth(0, 32)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 50)
	table.SetColumnWidth(3, 90)
	table.SetColumnWidth(4, 60)

	// Return the table.
	return table

}
