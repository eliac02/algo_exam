//Elia Cortesi 01911A

package main

import "fmt"

// bloccoOmog prints the sum of all the intensities of the tiles belonging to the block that the tile (x,y) belongs to, that are the same color of the tile (x,y)
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func bloccoOmog(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; !exists {
		fmt.Println("0")
		return
	}
	sum := 0
    block := make(map[piastrella]*properties)
    seen := make(map[piastrella]bool)
    block = trovaBloccoOmogeneo(p, tile, seen)
    for t := range block {
        if p.tiles[t].color == p.tiles[tile].color {
            sum += p.tiles[t].intensity
        }
    }
    fmt.Println(sum)
}
