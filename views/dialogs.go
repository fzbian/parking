package views

import (
	"fyne.io/fyne/v2/widget"
)

// AddVehicleDialog returns the form and the entries of the dialog to add a vehicle.
func AddVehicleDialog() (fyneForm *widget.Form, vehiclePlateEntry *widget.Entry, vehicleTypeEntry *widget.RadioGroup) {

	// Create the entries of the form.
	VehiclePlateEntry := widget.NewEntry()
	VehicleTypeEntry := widget.NewRadioGroup([]string{
		"Normal",
		"Vip",
		"Handicapped",
		"Emergency",
		"Provider"}, nil)

	// Create the form items.
	AddVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Plate", VehiclePlateEntry),
		widget.NewFormItem("Vehicle type", VehicleTypeEntry),
	}

	// Create the form.
	AddVehicleForm := &widget.Form{
		Items: AddVehicleFormItems,
	}

	// Return the form and the entries.
	return AddVehicleForm, VehiclePlateEntry, VehicleTypeEntry

}

// ExitVehicleDialog returns the form and the entry of the dialog to exit a vehicle.
func ExitVehicleDialog() (fyneForm *widget.Form, vehiclePlateExit *widget.Entry) {

	// Create the entry of the form.
	VehiclePlateExit := widget.NewEntry()

	// Create the form items.
	ExitVehicleFormItems := []*widget.FormItem{
		widget.NewFormItem("Plate", VehiclePlateExit),
	}

	// Create the form.
	ExitVehicleForm := &widget.Form{
		Items: ExitVehicleFormItems,
	}

	// Return the form and the entry.
	return ExitVehicleForm, VehiclePlateExit

}

// RecordsVehiclesTypeDialog returns the form and the entry of the dialog to get the records of a vehicle type.
func RecordsVehiclesTypeDialog() (fyneForm *widget.Form, VehicleType *widget.RadioGroup) {

	// Create the entry of the form.
	VehicleTypeEntry := widget.NewRadioGroup([]string{
		"Normal",
		"Vip",
		"Handicapped",
		"Emergency",
		"Provider"}, nil)

	// Create the form items.
	RecordsVehiclesTypeItems := []*widget.FormItem{
		widget.NewFormItem("Vehicle type", VehicleTypeEntry),
	}

	// Create the form.
	RecordsVehiclesTypeForm := &widget.Form{
		Items: RecordsVehiclesTypeItems,
	}

	// Return the form and the entry.
	return RecordsVehiclesTypeForm, VehicleTypeEntry

}

// RecordsZoneDialog returns the form and the entry of the dialog to get the records of a zone.
func RecordsZoneDialog() (fyneForm *widget.Form, zone *widget.RadioGroup) {

	// Create the entry of the form.
	ZoneEntry := widget.NewRadioGroup([]string{
		"A",
		"B",
		"C"}, nil)

	// Create the form items.
	RecordsZoneItems := []*widget.FormItem{
		widget.NewFormItem("Zone", ZoneEntry),
	}

	// Create the form.
	RecordsZoneForm := &widget.Form{
		Items: RecordsZoneItems,
	}

	// Return the form and the entry.
	return RecordsZoneForm, ZoneEntry

}
