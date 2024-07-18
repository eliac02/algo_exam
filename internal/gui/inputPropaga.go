package gui

import (
	"fmt"
	"strconv"
	"tiles/internal/algorithms"
	"tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// showPropagaDialogTakeParam takes from the user the input for Propaga
//
// @param ui The graphic interface
// @param p The system tiles-rules
func showPropagaDialogTakeParam(ui *models.UI, p models.Piano) {
	inputX := widget.NewEntry()
	inputX.SetPlaceHolder("X")
	inputY := widget.NewEntry()
	inputY.SetPlaceHolder("Y")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "X", Widget: inputX},
			{Text: "Y", Widget: inputY},
		},
		OnSubmit: func() {
			if err := validateTile(inputX.Text, inputY.Text); err != nil {
				inputX.SetText("")
				inputY.SetText("")
				errorLabel.SetText(err.Error())
				errorLabel.Show()
			} else {
				// Valori validi, procedere con le azioni desiderate
				errorLabel.SetText("Regola propagata")
				errorLabel.Show()
                x, err := strconv.Atoi(inputX.Text)
                if err != nil {
                    fmt.Println(err)
                }
                y, _ := strconv.Atoi(inputY.Text)
                if err != nil {
                    fmt.Println(err)
                }
                algorithms.Propaga(ui, p, x, y)
			}
		},
        SubmitText: "Propaga regola",
	}

	dialog := dialog.NewCustom("Seleziona piastrella su cui propagare", "Ok", container.NewVBox(
		form,
		errorLabel,
	), ui.Window)
	dialog.Resize(fyne.NewSize(200, 100))
	dialog.Show()
}

// validateTile checks the correctness of the user's input
//
// @param input1 The first input
// @param input2 The second input
// @return error The error that will be checked
func validateTile(input1, input2 string) error {
	if val, err := strconv.Atoi(input1); err != nil || val < 0 {
		return fmt.Errorf("X deve essere maggiore o uguale di zero")
	}
	if val, err := strconv.Atoi(input2); err != nil || val < 0 {
		return fmt.Errorf("Y deve essere maggiore o uguale di zero")
	}
	return nil
}
