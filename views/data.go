package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/fzbian/parking/controller"
	"github.com/fzbian/parking/models"
	"github.com/fzbian/parking/utils"
)

func GetTable() *widget.Table {
	var spots []models.Spot
	utils.Db.Find(&spots)

	var data [][]string

	for _, u := range spots {
		if u.InUse {
			plateNumber := controller.GetVehiclePlateNumberBySpotId(u.ID)
			data = append(data, []string{u.GetIDString(), u.Type, u.Zone, "Ocupado", plateNumber})
		} else {
			data = append(data, []string{u.GetIDString(), u.Type, u.Zone, "Libre", ""})
		}
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(data) + 1, len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("TableCell")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				switch i.Col {
				case 0:
					o.(*widget.Label).SetText("ID")
				case 1:
					o.(*widget.Label).SetText("Tipo")
				case 2:
					o.(*widget.Label).SetText("Zona")
				case 3:
					o.(*widget.Label).SetText("Estado")
				case 4:
					o.(*widget.Label).SetText("Placa")
				}
			} else {
				o.(*widget.Label).SetText(data[i.Row-1][i.Col])
			}
		},
	)

	table.SetColumnWidth(0, 32)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 50)
	table.SetColumnWidth(3, 90)
	table.SetColumnWidth(4, 60)

	return table
}
