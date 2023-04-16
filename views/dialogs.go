package views

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func AddVehicleDialog() (fyneForm *widget.Form, vehiclePlateEntry *widget.Entry, vehicleTypeEntry *widget.RadioGroup) {
	VehiclePlateEntry := widget.NewEntry()
	VehicleTypeEntry := widget.NewRadioGroup([]string{"NORMAL", "VIP", "DISCAPACITADO", "EMERGENCIA", "PROVEEDOR"}, nil)

	AddVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Placa", VehiclePlateEntry),
		widget.NewFormItem("Tipo de vehiculo", VehicleTypeEntry),
	}

	AddVehicleForm := &widget.Form{
		Items: AddVehicleFormItems,
	}
	return AddVehicleForm, VehiclePlateEntry, VehicleTypeEntry
}

func ExitVehicleDialog() (fyneForm *widget.Form, vehiclePlateExit *widget.Entry) {
	VehiclePlateExit := widget.NewEntry()

	ExitVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Placa", VehiclePlateExit),
	}

	ExitVehicleForm := &widget.Form{
		Items: ExitVehicleFormItems,
	}
	return ExitVehicleForm, VehiclePlateExit
}

func ExportRecordsDialog(window fyne.Window) {
	label := widget.NewLabel("Exportar registros")
	fyneForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Exportar", Widget: label},
		},
	}

	window.CenterOnScreen()
	fyneForm.Show()
}
