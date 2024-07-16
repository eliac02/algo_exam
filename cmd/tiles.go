// Elia Cortesi 01911A
package main

import (
	"tiles/internal/algorithms"
	"tiles/internal/gui"
	"tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Digital Tiles")
	p := algorithms.MakeSet()

    ui := &models.UI{App: a, Window: w}

    gui.InitializeGUI(ui, p)

    (*ui).Window.Resize(fyne.NewSize(800, 800))
	(*ui).Window.ShowAndRun()
}
