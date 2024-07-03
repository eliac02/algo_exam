// Elia Cortesi 01911A
package main

import "fmt"

// lung retrieve the lenght of the shortest path between two tiles
//
// @param p The system tiles-rules
// @param x1 x2 y1 y2 The coordinates of the two tiles
func lung(p piano, x1, y1, x2, y2 int) {
	start := piastrella{x: x1, y: y1}
	end := piastrella{x: x2, y: y2}

	// check if the tiles are existing and connected
	if _, exists := p.tiles[start]; !exists {
		return
	}
	if _, exists := p.tiles[end]; !exists {
		return
	}
	if !(p.Find(start) == p.Find(end)) {
		return
	}

	if start == end {
		// easy
		fmt.Println(1)
	} else {
		// calculate the shortest path
		lung := camminoMinimo(p, start, end)
		fmt.Println(lung)
	}
}
