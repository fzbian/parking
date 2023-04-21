package views

import (
	"fyne.io/fyne/v2/widget"
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

func RecordsVehiclesTypeDialog() (fyneForm *widget.Form, VehicleType *widget.RadioGroup) {
	VehicleTypeEntry := widget.NewRadioGroup([]string{"NORMAL", "VIP", "DISCAPACITADO", "EMERGENCIA", "PROVEEDOR"}, nil)

	RecordsVehiclesTypeItems := []*widget.FormItem{
		widget.NewFormItem("Tipo de vehiculo", VehicleTypeEntry),
	}

	RecordsVehiclesTypeForm := &widget.Form{
		Items: RecordsVehiclesTypeItems,
	}
	return RecordsVehiclesTypeForm, VehicleTypeEntry
}

func RecordsZoneDialog() (fyneForm *widget.Form, zone *widget.RadioGroup) {
	ZoneEntry := widget.NewRadioGroup([]string{"A", "B", "C"}, nil)

	RecordsZoneItems := []*widget.FormItem{
		widget.NewFormItem("Zona", ZoneEntry),
	}

	RecordsZoneForm := &widget.Form{
		Items: RecordsZoneItems,
	}
	return RecordsZoneForm, ZoneEntry
}
