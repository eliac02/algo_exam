// Elia Cortesi 01911A
package main

import "fmt"

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

	// calculate the shortest path
	lung := camminoMinimo(p, start, end)
	fmt.Println(lung)
}
