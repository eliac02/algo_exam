// Elia Cortesi 01911A
package main

import (
	"fmt"
)

// stato prints the color and the intensity of the tile (x,y)
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @return string The color of the tile
// @return int The intensity of the color of the tile
func stato(p piano, x, y int) (string, int) {
	tile := piastrella{x: x, y: y}
	// if the tile exists print its informations
	if _, exists := p.tiles[tile]; exists {
		fmt.Printf("%s %d\n", p.tiles[tile].color, p.tiles[tile].intensity)
		return p.tiles[tile].color, p.tiles[tile].intensity
	} else {
		return "", 0
	}
}
