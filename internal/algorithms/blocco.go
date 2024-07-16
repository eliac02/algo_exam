package algorithms

import (
    models "tiles/internal/models"
	"fmt"
)

// Blocco prints the sum of all the intensities of the tiles belonging to the block that the tile (x,y) belongs to
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func Blocco(p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}
	if _, exists := p.Tiles[tile]; exists {
		// find root of tile's block and get blockintensity
		root := p.Find(tile)
		fmt.Println(p.Tiles[root].BlockIntensity)
	} else {
		fmt.Println("0")
	}
}
