package main

import (
	"fmt"
)

func blocco(p piano, x, y int) {
    tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; exists {
		root := p.Find(tile)
		fmt.Println(p.tiles[root].blockIntensity)
	} else {
        fmt.Println("0")
    }
}
