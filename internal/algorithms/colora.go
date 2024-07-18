package algorithms

import (
	"fmt"
	models "tiles/internal/models"
	utils "tiles/internal/utils"
)

// Colora adds a tile to the system tiles-rules and colors it with a color and an intensity of color. If the tile exists already, colora updates it's color and intensity
//
// @param ui The graphic interface
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
// @param alpha The color of the tile
// @param i The intensity of the color
func Colora(ui *models.UI, p models.Piano, x, y int, hex string, i int) {
	tile := models.Piastrella{X: x, Y: y}

	// if intensity = 0 then ignore it
	if i <= 0 {
		return
	} else {
		// create the tile and add it to the system
		p.Add(tile, hex, i)
	}

	// if it has adjacents execute the Unions
	adiacenti := utils.GetAdiacenti(x, y)
	for _, adj := range adiacenti {
		if _, exists := p.Tiles[adj]; exists {
			p.Union(tile, adj)
		}
	}

	col, err := utils.ParseHexColor(hex)
	if err != nil {
		fmt.Println(err)
	}
	ui.UpdateCell(x, 14-y, col)
}
