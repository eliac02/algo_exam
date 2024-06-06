package main

func propaga(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	colorCount := make(map[string]int)
	adiacenti := getAdiacenti(x, y)
	for _, adj := range adiacenti {
		if props, exists := p.tiles[adj]; exists {
			colorCount[props.color]++
		}
	}

	for _, reg := range *p.rules {
		if ruleOk(reg, colorCount) {
			if _, exists := p.tiles[tile]; exists {
				p.tiles[tile].color = reg.color
				break
			} else {
                colora(p, x, y, reg.color, 1)
            }
		}
	}
}
