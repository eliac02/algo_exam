// Elia Cortesi 01911A
package main

import (
    "tiles/internal/models"
    "tiles/internal/algorithms"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)



func main() {
	InitializeGUI()

	// p := makeSet()
	// grid := newDynamicGrid()

	button := widget.NewButton("Aggiungi elemento", showDialog)

	w.SetContent(container.NewVBox(button))

	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		esegui(p, scanner.Text())
	}*/

	w.ShowAndRun()
}
