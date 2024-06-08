package main

func getAdiacenti(x, y int) [8]piastrella {
	adjs := [8]piastrella{
		{x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	return adjs
}

func ruleOk(reg rule, colorCount map[string]int) bool {
	for _, rs := range reg.ruleset {
		if colorCount[rs.color] < rs.count {
			return false
		}
	}
	return true
}

func esploraViciniOmog(p piano, x, y int, seen map[piastrella]bool, sum *int) {
	tile := piastrella{x: x, y: y}
	coda := []piastrella{tile}
	*sum += p.tiles[tile].intensity

	for len(coda) > 0 {
		current := coda[0]
		coda = coda[1:]
		adiacenti := getAdiacenti(current.x, current.y)

		for _, adj := range adiacenti {
			if _, exists := p.tiles[adj]; exists && !seen[adj] {
				seen[adj] = true
				if p.tiles[adj].color == p.tiles[tile].color {
					*sum += p.tiles[adj].intensity
				}
				coda = append(coda, adj)
			}
		}
	}
}

func esploraBlocco(p piano, start piastrella, seen map[piastrella]bool) map[piastrella]*properties {
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
				if p.Find(adj) == p.Find(start) {
					pila = append(pila, adj)
				}
			}
		}
	}
	return block
}

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

