package main

import "fmt"

func bloccoOmog(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; !exists {
		fmt.Println("0")
		return
	}
    root := p.Find(tile)
	sum := 0
    block := make(map[piastrella]*properties)
    seen := make(map[piastrella]bool)
    block = trovaBlocco(p, root, seen)
    for t := range block {
        if p.tiles[t].color == p.tiles[tile].color {
            sum += p.tiles[t].intensity
        }
    }
    fmt.Println(sum)
}
