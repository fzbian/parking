package views

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
)

func GetData() [][]string {
	var spots []models.Spot
	utils.Db.Find(&spots)

	var data [][]string

	for _, u := range spots {
		var text string
		if u.InUse {
			text = "Ocupado"
		} else {
			text = "Libre"
		}
		data = append(data, []string{fmt.Sprint(u.ID), u.Type, u.Zone, text})
	}

	return data
}

func GetTable(data [][]string) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(cell widget.TableCellID, cellView fyne.CanvasObject) {
			cellView.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 90)
	table.SetColumnWidth(2, 50)
	table.SetColumnWidth(3, 50)

	return table
}

func GetVehiclesTable() *widget.Table {
	data := GetData()
	table := GetTable(data)
	return table
}

func RefreshTable() {
	VehicleTable := GetVehiclesTable()
	VehicleTable.Refresh()
}
