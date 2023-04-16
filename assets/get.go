package assets

import "fyne.io/fyne"

func GetIcon() fyne.Resource {
	LogoResource, err := fyne.LoadResourceFromPath("assets/icon.png")
	if err != nil {
		panic(err)
	}

	return LogoResource
}
