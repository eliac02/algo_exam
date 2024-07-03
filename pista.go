// Elia Cortesi 01911A
package main

import (
	"fmt"
	"strings"
)

// pista prints the sequence of tiles following the directions of the string s
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @param s The list of directions
func pista(p piano, x, y int, s string) {
	tile := piastrella{x: x, y: y}

	if _, exists := p.tiles[tile]; !exists {
		return
	}

	directions := strings.Split(s, " ")
	listOfDirections := directions[3]

	// if the directions are valid print the tiles
	flag, sequence := verificaPista(p, x, y, listOfDirections)
	if flag {
		fmt.Println("[")
		fmt.Printf("%d %d %s %d\n", tile.x, tile.y, p.tiles[tile].color, p.tiles[tile].intensity)
		for _, el := range sequence {
			fmt.Printf("%d %d %s %d\n", el.x, el.y, p.tiles[el].color, p.tiles[el].intensity)
		}
		fmt.Println("]")
	} else {
		return
	}
}
