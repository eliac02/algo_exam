package algorithms

import (
	"fmt"
	models "tiles/internal/models"
)

// Stato prints the color and the intensity of the tile (x,y)
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @return string The color of the tile
// @return int The intensity of the color of the tile
func Stato(p models.Piano, x, y int) (string, int) {
	tile := models.Piastrella{X: x, Y: y}
	// if the tile exists print its informations
	if _, exists := p.Tiles[tile]; exists {
		fmt.Printf("%s %d\n", p.Tiles[tile].Color, p.Tiles[tile].Intensity)
		return p.Tiles[tile].Color, p.Tiles[tile].Intensity
	} else {
		return "", 0
	}
}
