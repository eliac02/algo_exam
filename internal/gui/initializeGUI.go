package gui

import (
	"image/color"
	models "tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// initGrid initializes the graphic grid
//
// @param ui The graphic interface
func initGrid(ui *models.UI) {
	ui.Rects = make([][]*canvas.Rectangle, ui.Columns)
	for i := range ui.Rects {
		ui.Rects[i] = make([]*canvas.Rectangle, ui.Rows)
		for j := range ui.Rects[i] {
			ui.Rects[i][j] = canvas.NewRectangle(color.Transparent)
			ui.Rects[i][j].StrokeColor = color.Gray{Y: 0xAA}
			ui.Rects[i][j].StrokeWidth = 1
			ui.Rects[i][j].SetMinSize(fyne.Size{Width: 47, Height: 40})
		}
	}
	ui.Grid = container.NewGridWithColumns(ui.Columns)

	for _, row := range ui.Rects {
		for _, rect := range row {
			ui.Grid.Add(rect)
		}
	}
}

// initButtons initializes the buttons of the functionalities
//
// @param ui The graphic interface
// @param p The system tiles-rules
func initButtons(ui *models.UI, p models.Piano) {
	ui.Buttons = container.NewGridWithColumns(1)

	buttons := make([]*widget.Button, 5)

	coloraButton := widget.NewButton("Colora", func() {
		showColoraDialogTakeParam(ui, p)
	})
	buttons[0] = coloraButton

	regolaButton := widget.NewButton("Regola", func() {
		showRegolaDialogTakeParam(ui, p)
	})
	buttons[1] = regolaButton

	printButton := widget.NewButton("Stampa", func() {
		showStampaDialog(ui, p)
	})
	buttons[2] = printButton

	propagaButton := widget.NewButton("Propaga", func() {
		showPropagaDialogTakeParam(ui, p)
	})
	buttons[3] = propagaButton

    spegniButton := widget.NewButton("Spegni", func() {
        showSpegniDialogTakeParam(ui, p)
    })
    buttons[4] = spegniButton

	for _, but := range buttons {
		ui.Buttons.Add(but)
	}
}

// initList initialize the list of rules
//
// @param ui The graphic interface
func initList(ui *models.UI) {
    ui.List = container.NewGridWithColumns(1)
    rules := make([]*widget.Label, 0)
    ui.Rules = &rules
}

// initializeGUI initializes the whole graphic interface
//
// @param ui The graphic interface
// @param p The system tiles-rules
func InitializeGUI(ui *models.UI, p models.Piano) {
	ui.Columns = 15
	ui.Rows = 15

	initGrid(ui)

	initButtons(ui, p)

	initList(ui)

	ui.Window.SetContent(container.NewHBox(
		ui.Grid,
		ui.Buttons,
	))
}
