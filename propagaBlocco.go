//Elia Cortesi 01911A

package main

// propagaBlocco applies the first rule available to each of the tiles of the block that the tile (x,y) belongs to
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
func propagaBlocco(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
    root := p.Find(tile)
	seen := make(map[piastrella]bool)
	block := trovaBlocco(p, root, seen, func(piastrella) bool {
        return true
    })

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

		for index, reg := range *p.rules {
			if ruleOk(reg, colorCount) {
				p.tiles[t].color = reg.color
				(*p.rules)[index].usage++
				break
			}
		}

	}
}
