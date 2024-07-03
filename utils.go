// Elia Cortesi 01911A
package main

import "strings"

// getAdiacenti returns the list of the tiles adjacents to the tile (x,y)
//
// @param x y The coordinates of the tile
// @return [8]piastrella The list of the tiles adjacents to the tile (x,y)
func getAdiacenti(x, y int) [8]piastrella {
	// retrieve the adjacents tiles
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
	// check if the rule can be applied
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
// @param filtro Lambda function that checks a given requirement
// @return map[piastrella]*properties The block that the tile start belongs to
func trovaBlocco(p piano, start piastrella, seen map[piastrella]bool, filtro func(p piastrella) bool) map[piastrella]*properties {
	block := make(map[piastrella]*properties)
	pila := []piastrella{start}

	// use a DFS to retrieve the block's tiles
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
				// use a lambda function as a filter
				if filtro(adj) {
					pila = append(pila, adj)
				}
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
	block := trovaBlocco(p, root, seen, func(piastrella) bool {
		return true
	})

	// change the root to the first adjacent tile found
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

// verificaPista checks if the sequence of directions is valid
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @param directions The list of directions
// @return bool True if the sequence is valid, False otherwise
func verificaPista(p piano, x, y int, s string) (bool, []piastrella) {
	sequence := make([]piastrella, 0)
	// assign to each direction an index
	corresponding := map[string]int{
		"NN": 0,
		"NO": 1,
		"OO": 2,
		"SO": 3,
		"SS": 4,
		"SE": 5,
		"EE": 6,
		"NE": 7,
	}

	directions := strings.Split(s, ",")

	// check one-by-one if all the directions are valid
	for _, dir := range directions {
		adiacenti := getAdiacenti(x, y)
		tile := adiacenti[corresponding[dir]]
		if _, exists := p.tiles[tile]; !exists {
			return false, []piastrella{}
		}
		sequence = append(sequence, tile)
		x, y = tile.x, tile.y
	}
	return true, sequence
}

// camminoMinimo calculate the shortest path from start to end
//
// @param p The system tiles-rules
// @param start The starting tile
// @param end The ending tile
// @return int The lenght of the shortest path between @start and @end
func camminoMinimo(p piano, start, end piastrella) int {
	coda := []piastrella{start}
	distanze := make(map[piastrella]int)
	distanze[start] = 1

	// use a BFS to calculate the shortest path
	for len(coda) > 0 {
		current := coda[0]
		coda = coda[1:]

		if current == end {
			return distanze[current]
		}

		adiacenti := getAdiacenti(current.x, current.y)
		for _, adj := range adiacenti {
			if _, exists := p.tiles[adj]; exists {
				if _, visited := distanze[adj]; !visited {
					coda = append(coda, adj)
					distanze[adj] = distanze[current] + 1
				}
			}
		}
	}

	return 0
}
