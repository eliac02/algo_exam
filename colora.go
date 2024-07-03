// Elia Cortesi 01911A
package main

// colora adds a tile to the system tiles-rules and colors it with a color and an intensity of color. If the tile exists already, colora updates it's color and intensity
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
// @param alpha The color of the tile
// @param i The intensity of the color
func colora(p piano, x, y int, alpha string, i int) {
	tile := piastrella{x: x, y: y}

	// if intensity = 0 then turn off tile
	if i <= 0 {
		spegni(p, x, y)
		return
	} else {
		// create the tile and add it to the system
		p.Add(tile, alpha, i)
	}

	// if it has adjacents execute the Unions
	adiacenti := getAdiacenti(x, y)
	for _, adj := range adiacenti {
		if _, exists := p.tiles[adj]; exists {
			p.Union(tile, adj)
		}
	}
}
