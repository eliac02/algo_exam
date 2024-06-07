package main

import "fmt"

func bloccoOmog(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; !exists {
		fmt.Println("0")
		return
	}
	somma := 0
	var sum *int = &somma
	seen := make(map[piastrella]bool)
	seen[piastrella{x: x, y: y}] = true
	esploraViciniOmog(p, x, y, seen, sum)
    fmt.Println(somma)
}
