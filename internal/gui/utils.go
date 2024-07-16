package gui

import (
	"fmt"
	"image/color"
	"strconv"
	"tiles/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// checkGridSize checks if a newly created tile's coordinates are bugger than the sizes of the checkGridSize
//
// @param ui The graphic interface
// @param tile The newly created tile
func checkGridSize(ui *models.UI, x, y int) {
	// if necessary fix the size of the grid
	if x > ui.Columns {
		ui.Columns = x
	}
	if y > ui.Rows {
		ui.Rows = y
	}
}

func initGrid(ui *models.UI) {
	ui.Rects = make([][]*canvas.Rectangle, ui.Columns)
	for i := range ui.Rects {
		ui.Rects[i] = make([]*canvas.Rectangle, ui.Rows)
		for j := range ui.Rects[i] {
			ui.Rects[i][j] = canvas.NewRectangle(color.Transparent)
			ui.Rects[i][j].StrokeColor = color.Gray{Y: 0xAA}
			ui.Rects[i][j].StrokeWidth = 1
            ui.Rects[i][j].SetMinSize(fyne.Size{Width: 60,Height: 40})
		}
	}
	ui.Grid = container.NewGridWithColumns(ui.Columns)

	for _, row := range ui.Rects {
		for _, rect := range row {
			ui.Grid.Add(rect)
		}
	}
}

// UpdateGrid updates the grid when a new tile is coloured
//
// @param ui The graphic interface
// @param p The system tiles-rules
func UpdateCell(ui *models.UI, x, y int, hex color.Color) {
	ui.Rects[y][x].FillColor = hex
	ui.Rects[y][x].Refresh()
}

// parseHexColor parses the hexadecimal color of a tile in a color.Color object
//
// @param s The hex string
// @return color.Color The object
// @error An error that can occur
func parseHexColor(hex string) (color.Color, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return nil, fmt.Errorf("Formato colore non valido: %s", hex)
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return nil, err
	}
	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return nil, err
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}, nil
}
