package algorithms

import (
	"fmt"
	models "tiles/internal/models"
	utils "tiles/internal/utils"
)

// Lung retrieve the lenght of the shortest path between two tiles
//
// @param p The system tiles-rules
// @param x1 x2 y1 y2 The coordinates of the two tiles
func Lung(p models.Piano, x1, y1, x2, y2 int) {
	start := models.Piastrella{X: x1, Y: y1}
	end := models.Piastrella{X: x2, Y: y2}

	// check if the tiles are existing and connected
	if _, exists := p.Tiles[start]; !exists {
		return
	}
	if _, exists := p.Tiles[end]; !exists {
		return
	}
	if !(p.Find(start) == p.Find(end)) {
		return
	}

	if start == end {
		// easy
		fmt.Println(1)
	} else {
		// calculate the shortest path
		lung := utils.CamminoMinimo(p, start, end)
		fmt.Println(lung)
	}
}
