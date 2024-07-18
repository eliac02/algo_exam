package gui

import (
	"fmt"
	"strconv"
	"strings"
	"tiles/internal/algorithms"
	"tiles/internal/models"
	"tiles/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// showRegolaDialogTakeParam takes from the user the input for Regola
//
// @param ui The graphical interface
// @param p The system tiles-rules
func showRegolaDialogTakeParam(ui *models.UI, p models.Piano) {
	input := widget.NewEntry()
	input.SetPlaceHolder("colore finale(hex) numero colore(hex) numero colore(hex)...")
	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Regola", Widget: input},
		},
		OnSubmit: func() {
			if err := validateRule(input.Text); err != nil {
				input.SetText("")
				errorLabel.SetText(err.Error())
				errorLabel.Show()
			} else {
				// Valori validi, procedere con le azioni desiderate
				errorLabel.SetText("Regola inserita")
				errorLabel.Show()
				algorithms.Regola(p, input.Text)
				utils.AggiungiRegola(ui, p)
			}
		},
		SubmitText: "Inserisci regola",
	}

	dialog := dialog.NewCustom("Inserisci regola", "Ok", container.NewVBox(
		form,
		errorLabel,
	), ui.Window)
	dialog.Resize(fyne.NewSize(550, 200))
	dialog.Show()
}

// validateRule checks the correctness of the user's input
//
// @param rule The input
// @return error The error that will be checked
func validateRule(rule string) error {
	controlMap := make(map[string]bool)
	ruleSplitted := strings.Split(rule, " ")
	counterTiles := 0

	if !utils.IsValidHexColor(ruleSplitted[0]) {
		return fmt.Errorf("Il colore finale non è nel formato corretto")
	}

	for i := 1; i < len(ruleSplitted); i++ {
		if i%2 != 0 {
			num, _ := strconv.Atoi(ruleSplitted[i])
			counterTiles += num
			if counterTiles > 8 {
				return fmt.Errorf("La somma dei numeri deve essere al massimo 8")
			}
		} else {
			if !utils.IsValidHexColor(ruleSplitted[i]) || controlMap[ruleSplitted[i]] {
				return fmt.Errorf("Uno dei colori non è nel formato corretto o è già stato inserito")
			} else {
				controlMap[ruleSplitted[i]] = true
			}
		}
	}
	return nil
}
