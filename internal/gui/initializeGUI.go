package gui

import (
	"tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitializeGUI(a fyne.App, w fyne.Window) *models.UI {
    buttons := make(map[models.Piastrella]widget.Button)
	ui := &models.UI{
		App:    a,
		Window: w,
		Buttons: &buttons,
    }

	ui.Window.SetContent(container.NewHBox(
	// griglia,
	// bottoni,
	))

	return ui
}
