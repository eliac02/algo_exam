package algorithms

import (
	"fmt"
	models "tiles/internal/models"
	utils "tiles/internal/utils"
)

// Propaga applies the first available rule to the tile (x,y)
//
// @param ui The graphic interface
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func Propaga(ui *models.UI, p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}

	// retrieve the adjacent colors
	colorCount := make(map[string]int)
	adiacenti := utils.GetAdiacenti(x, y)
	for _, adj := range adiacenti {
		if props, exists := p.Tiles[adj]; exists {
			colorCount[props.Color]++
		}
	}

	// check if a rule can be applied
	for index, reg := range *p.Rules {
		if utils.RuleOk(reg, colorCount) {
			if _, exists := p.Tiles[tile]; exists {
				p.Tiles[tile].Color = reg.Color
				hex, err := utils.ParseHexColor(reg.Color)
				if err != nil {
					fmt.Println(err)
				}
				ui.UpdateCell(x, 14-y, hex)
			} else {
				Colora(ui, p, x, y, reg.Color, 1)
			}
            ui.UpdateRule(reg)
			(*p.Rules)[index].Usage++
            fmt.Println((*p.Rules)[index].Usage)
			break
		}
	}
}
