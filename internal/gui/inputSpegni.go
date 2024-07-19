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

// showSpegniDialogTakeParam takes from the user the input for Spegni
//
// @param win The graphic interface
// @param p The system tiles-rules
func showSpegniDialogTakeParam(ui *models.UI, p models.Piano) {
	inputX := widget.NewEntry()
	inputX.SetPlaceHolder("X")
	inputY := widget.NewEntry()
	inputY.SetPlaceHolder("Y")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Ascissa", Widget: inputX},
			{Text: "Ordinata", Widget: inputY},
		},
		OnSubmit: func() {
			errorLabel.Hide()
			if err := validateInput(inputX.Text, inputY.Text); err != nil {
				inputX.SetText("")
				inputY.SetText("")
				errorLabel.SetText(err.Error())
				errorLabel.Show()
			} else {
				// Valori validi, procedere con le azioni desiderate
				x, err := strconv.Atoi(inputX.Text)
				if err != nil {
					fmt.Println(err)
				}
				y, err := strconv.Atoi(inputY.Text)
				if err != nil {
					fmt.Println(err)
				}
				algorithms.Spegni(ui, p, x, y)
			}
		},
		SubmitText: "Spegni",
	}

	dialog := dialog.NewCustom("Inserisci Parametri", "Ok", container.NewVBox(
		form,
		errorLabel,
	), ui.Window)
	dialog.Resize(fyne.NewSize(250, 250))
	dialog.Show()
}

// validateInput checks the correctness of the user's input
//
// @param1 input1 The first input
// @param2 input2 The second input
// @return error The error that will be checked
func validateInput(input1, input2 string) error {
	if val, err := strconv.Atoi(input1); err != nil || val < 0 {
		return fmt.Errorf("X deve essere maggiore o uguale di zero")
	}
	if val, err := strconv.Atoi(input2); err != nil || val < 0 {
		return fmt.Errorf("Y deve essere maggiore o uguale di zero")
	}
	return nil
}
