package algorithms

import models "tiles/internal/models"

// Propaga applies the first available rule to the tile (x,y)
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func Propaga(p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}

	// retrieve the adjacent colors
	colorCount := make(map[string]int)
	adiacenti := getAdiacenti(x, y)
	for _, adj := range adiacenti {
		if props, exists := p.Tiles[adj]; exists {
			colorCount[props.Color]++
		}
	}

	// check if a rule can be applied
	for index, reg := range *p.Rules {
		if ruleOk(reg, colorCount) {
			if _, exists := p.Tiles[tile]; exists {
				p.Tiles[tile].Color = reg.Color
			} else {
				Colora(p, x, y, reg.Color, 1)
			}
			(*p.Rules)[index].Usage++
			break
		}
	}
}
