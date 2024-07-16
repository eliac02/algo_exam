package gui

import (
	"fmt"
	"regexp"
	"strconv"
	"tiles/internal/algorithms"
	"tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// showColoraDialogTakeParam takes from the user the input for Colora
//
// @param win The window where the dialog will be displayed
// @return x y The coordinates of the tile
// @return hex The hexadecimal color of the tile
// @intensity The intensity of the color of the tile
func showColoraDialogTakeParam(win fyne.Window, p models.Piano) {
	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Ascissa")
	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Ordinata")
	input3 := widget.NewEntry()
	input3.SetPlaceHolder("Colore (es: #rrggbb)")
	input4 := widget.NewEntry()
	input4.SetPlaceHolder("Intensità colore")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Ascissa", Widget: input1},
			{Text: "Ordinata", Widget: input2},
			{Text: "Colore", Widget: input3},
			{Text: "Intensità", Widget: input4},
		},
		OnSubmit: func() {
			errorLabel.Hide()
			if err := validateInputs(input1.Text, input2.Text, input3.Text, input4.Text); err != nil {
				errorLabel.SetText(err.Error())
				errorLabel.Show()
			} else {
				// Valori validi, procedere con le azioni desiderate
				x, err := strconv.Atoi(input1.Text)
				if err != nil {
					fmt.Println(err)
				}
				y, err := strconv.Atoi(input2.Text)
				if err != nil {
					fmt.Println(err)
				}
				i, err := strconv.Atoi(input4.Text)
				if err != nil {
					fmt.Println(err)
				}

				algorithms.Colora(p, x, y, input3.Text, i)
                fmt.Println(p.Tiles[models.Piastrella{X: x, Y: y}].Color)
			}
		},
	}

	dialog.ShowCustom("Inserisci Parametri", "Ok", container.NewVBox(
		form,
		errorLabel,
	), win)
}

// validateInput checks the correctness of the user's input
//
// @param input1 The first input
// @param2 input2 The second input
// @param3 input3 The third input
// @param4 input4 The fourth input
// @return error The error that will be checked
func validateInputs(input1, input2, input3, input4 string) error {
	if val, err := strconv.Atoi(input1); err != nil || val < 0 {
		return fmt.Errorf("Deve essere maggiore o uguale di zero")
	}
	if val, err := strconv.Atoi(input2); err != nil || val < 0 {
		return fmt.Errorf("Deve essere maggiore o uguale di zero")
	}
	if !isValidHexColor(input3) {
		return fmt.Errorf("Deve essere un valore esadecimale valido (#rrggbb)")
	}
	if val, err := strconv.Atoi(input4); err != nil || val <= 0 {
		return fmt.Errorf("Deve essere maggiore di zero")
	}
	return nil
}

// isValidHexColor checks if the hexadecimal color is correct
//
// @param s The hexadecimal string
// @return bool Returns True if the color is corretc, False otherwise
func isValidHexColor(s string) bool {
	re := regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	return re.MatchString(s)
}
