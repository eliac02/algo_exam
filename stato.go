package main

import (
	"fmt"
)

func stato(p piano, x, y int) (string, int) {
	if _, exists := p.tiles[piastrella{x, y}]; exists {
		fmt.Printf("%s %d\n", p.tiles[piastrella{x, y}].color, p.tiles[piastrella{x, y}].intensity)
		return p.tiles[piastrella{x, y}].color, p.tiles[piastrella{x, y}].intensity
	} else {
		return "null", 0
	}
}
