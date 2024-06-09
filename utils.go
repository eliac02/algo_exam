//Elia Cortesi 01911A

package main

// getAdiacenti returns the list of the tiles adjacents to the tile (x,y)
//
// @param x y The coordinates of the tile
// @return [8]piastrella The list of the tiles adjacents to the tile (x,y)
func getAdiacenti(x, y int) [8]piastrella {
	adjs := [8]piastrella{
		{x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	return adjs
}

// ruleOk checks if a rule can be applied to a tile
// 
// @param reg The rule to check the validity of
// @colorCount A map that contains the color of the tiles adjacents to a tile
// @return bool True if the rule can be applied, False otherwise
func ruleOk(reg rule, colorCount map[string]int) bool {
	for _, rs := range reg.ruleset {
		if colorCount[rs.color] < rs.count {
			return false
		}
	}
	return true
}

// trovaBlocco finds all the tiles belonging to the block that the tile start belongs to
// 
// @param p The system tiles-rules
// @param start The tile which start the search from
// @return map[piastrella]*properties The block that the tile start belongs to
func trovaBlocco(p piano, start piastrella, seen map[piastrella]bool) map[piastrella]*properties {
	block := make(map[piastrella]*properties)
	pila := []piastrella{start}

	for len(pila) > 0 {
		current := pila[len(pila)-1]
		pila = pila[:len(pila)-1]

        if seen[current] {
            continue
        }

        seen[current] = true
        block[current] = p.tiles[current]

		adiacenti := getAdiacenti(current.x, current.y)
		for _, adj := range adiacenti {
			if _, exists := p.tiles[adj]; exists && !seen[adj] {
					pila = append(pila, adj)
				}
			}
		}
	return block
}

// cambiaRadice change the root of a block
//
// @param p The system tiles-rules
// @param root The root that needs to be changes
func cambiaRadice(p piano, root piastrella) {
	adiacenti := getAdiacenti(root.x, root.y)
	seen := make(map[piastrella]bool)
	block := trovaBlocco(p, root, seen)
	for _, adj := range adiacenti {
		if _, exists := p.tiles[adj]; exists {
			p.tiles[root].blockIntensity -= p.tiles[root].intensity
			p.tiles[root].rank--
			p.tiles[adj].blockIntensity = p.tiles[root].blockIntensity
			p.tiles[adj].rank = p.tiles[root].rank
			p.tiles[adj].parent = adj
			for t := range block {
				p.tiles[t].parent = adj
			}
			return
		}
	}
}

