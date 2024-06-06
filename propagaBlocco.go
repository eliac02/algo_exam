package main

func propagaBlocco(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	root := p.Find(tile)
	seen := make(map[piastrella]bool)
	block := esploraBlocco(p, root, seen)

	originalBlock := make(map[piastrella]*properties)
	for t, props := range block {
		originalBlock[t] = &properties{
			color:          props.color,
			intensity:      props.intensity,
			parent:         props.parent,
			rank:           props.rank,
			blockIntensity: props.blockIntensity,
		}
	}

    for t := range block {
		colorCount := make(map[string]int)
		adiacenti := getAdiacenti(t.x, t.y)
		for _, adj := range adiacenti {
			if props, exists := originalBlock[adj]; exists {
				colorCount[props.color]++
			}
		}

		for _, reg := range *p.rules {
			if ruleOk(reg, colorCount) {
				p.tiles[tile].color = reg.color
				break
			}
		}

	}
}
