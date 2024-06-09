// Elia Cortesi 01911A
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

	// retrieve the color of the tile
	color := p.tiles[tile].color
	sum := 0
	block := make(map[piastrella]*properties)
	seen := make(map[piastrella]bool)

    // use the color of tile as a condition in the filter lambda function
	block = trovaBlocco(p, tile, seen, func(adj piastrella) bool {
		return p.tiles[adj].color == color
	})

    //calculate total intensity
	for t := range block {
		sum += p.tiles[t].intensity
	}
	fmt.Println(sum)
}
