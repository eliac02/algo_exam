package gui

import (
	"tiles/internal/algorithms"
	"tiles/internal/models"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// showStampaDialog let the user see the lust of rules
//
// @param ui The graphic interface
// @param p The system tiles-rules
func showStampaDialog(ui *models.UI, p models.Piano) {
	windowRule := ui.App.NewWindow("Lista regole")

	sortButton := widget.NewButton("Ordina", func() {
		algorithms.Ordina(p)
		ui.List.Refresh()
	})

	closeButton := widget.NewButton("Chiudi", func() {
		windowRule.Close()
	})

	windowRule.SetContent(container.NewVBox(
		widget.NewLabel("Lista di regole"),
		ui.List,
		sortButton,
		closeButton,
	))

	windowRule.Show()
}
