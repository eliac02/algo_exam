package algorithms

import (
	"fmt"
	"strings"
	models "tiles/internal/models"
)

// Pista prints the sequence of tiles following the directions of the string s
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @param s The list of directions
func Pista(p models.Piano, x, y int, s string) {
	tile := models.Piastrella{X: x, Y: y}

	if _, exists := p.Tiles[tile]; !exists {
		return
	}

	directions := strings.Split(s, " ")
	listOfDirections := directions[3]

	// if the directions are valid print the tiles
	flag, sequence := verificaPista(p, x, y, listOfDirections)
	if flag {
		fmt.Println("[")
		fmt.Printf("%d %d %s %d\n", tile.X, tile.Y, p.Tiles[tile].Color, p.Tiles[tile].Intensity)
		for _, el := range sequence {
			fmt.Printf("%d %d %s %d\n", el.X, el.Y, p.Tiles[el].Color, p.Tiles[el].Intensity)
		}
		fmt.Println("]")
	} else {
		return
	}
}
