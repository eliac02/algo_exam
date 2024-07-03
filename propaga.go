// Elia Cortesi 01911A
package main

// propaga applies the first available rule to the tile (x,y)
//
// @param p The system of tiles-rules
// @param x y The coordinates of the tile
func propaga(p piano, x, y int) {
	tile := piastrella{x: x, y: y}

	// retrieve the adjacent colors
	colorCount := make(map[string]int)
	adiacenti := getAdiacenti(x, y)
	for _, adj := range adiacenti {
		if props, exists := p.tiles[adj]; exists {
			colorCount[props.color]++
		}
	}

	// check if a rule can be applied
	for index, reg := range *p.rules {
		if ruleOk(reg, colorCount) {
			if _, exists := p.tiles[tile]; exists {
				p.tiles[tile].color = reg.color
			} else {
				colora(p, x, y, reg.color, 1)
			}
			(*p.rules)[index].usage++
			break
		}
	}
}
