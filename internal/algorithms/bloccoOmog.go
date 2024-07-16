package algorithms

import (
    models "tiles/internal/models"
    "fmt"
)

// BloccoOmog prints the sum of all the intensities of the tiles belonging to the block that the tile (x,y) belongs to, that are the same color of the tile (x,y)
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func BloccoOmog(p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}
	if _, exists := p.Tiles[tile]; !exists {
		fmt.Println("0")
		return
	}

	// retrieve the color of the tile
	color := p.Tiles[tile].Color
	sum := 0
	block := make(map[models.Piastrella]*models.Properties)
	seen := make(map[models.Piastrella]bool)

	// use the color of tile as a condition in the filter lambda function
	block = trovaBlocco(p, tile, seen, func(adj models.Piastrella) bool {
		return p.Tiles[adj].Color == color
	})

	// calculate total intensity
	for t := range block {
		sum += p.Tiles[t].Intensity
	}
	fmt.Println(sum)
}
