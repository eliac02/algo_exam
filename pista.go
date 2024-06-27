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

	// if the directions are valid print the tiles
	flag, sequence := verificaPista(p, x, y, directions)
	if flag {
		fmt.Println("(")
		for _, el := range sequence {
			fmt.Printf("%s %d\n", p.tiles[el].color, p.tiles[el].intensity)
		}
		fmt.Println(")")
	} else {
		return
	}
}
