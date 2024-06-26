// Elia Cortesi 01911A
package main

// spegni removes from the system the tile (x,y)
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
func spegni(p piano, x, y int) {
	tile := piastrella{x: x, y: y}

	if _, exists := p.tiles[tile]; !exists {
		return
	}

    root := p.Find(tile)

    //if the tile is the root, change the root
	if tile == root {
		cambiaRadice(p, root)
	}

    //retrieve the new blocks that may have been created
	adiacenti := getAdiacenti(x, y)
	seen := make(map[piastrella]bool)
	seen[tile] = true
	otherBlocks := make([]map[piastrella]*properties, 0)

	for _, adj := range adiacenti {
		if _, exists := p.tiles[adj]; exists && !seen[adj] {
			newBlock := trovaBlocco(p, adj, seen, func(piastrella) bool {
				return true
			})
			otherBlocks = append(otherBlocks, newBlock)
		}
	}

    //fix the information of the new blocks
	if len(otherBlocks) > 0 {
		for _, block := range otherBlocks {
			var newRoot piastrella
			totalIntensity := 0
			maxRank := 0

			for t := range block {
				totalIntensity += p.tiles[t].intensity
				maxRank++
				if newRoot == (piastrella{}) {
					newRoot = t
				}
				p.tiles[t].parent = newRoot
				if p.tiles[t].blockIntensity > p.tiles[t].intensity {
					p.tiles[t].blockIntensity = p.tiles[t].intensity
				}
				if p.tiles[t].rank > 0 {
					p.tiles[t].rank = 0
				}
			}

			p.tiles[newRoot].blockIntensity = totalIntensity
			p.tiles[newRoot].rank = maxRank
		}
	}

    //ddelete the tile from the system
	delete(p.tiles, tile)
}
