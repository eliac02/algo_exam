// Elia Cortesi 01911A

package main

import (
	"fmt"
)

// blocco prints the sum of all the intensities of the tiles belonging to the block that the tile (x,y) belongs to
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func blocco(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; exists {
		// find root of tile's block and get blockintensity
		root := p.Find(tile)
		fmt.Println(p.tiles[root].blockIntensity)
	} else {
		fmt.Println("0")
	}
}
