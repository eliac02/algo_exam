package main

func spegni(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	if _, exists := p.tiles[tile]; !exists {
		return
	}

	root := p.Find(tile)
    originalBlockIntensity := p.tiles[root].blockIntensity
    
	p.tiles[root].blockIntensity -= p.tiles[tile].intensity

	delete(p.tiles, tile)

	adiacenti := getAdiacenti(x, y)
	var otherBlocks []map[piastrella]*properties
	seen := make(map[piastrella]bool)
	seen[tile] = true

	for _, adj := range adiacenti {
		if _, exists := p.tiles[adj]; exists && !seen[adj] {
			if p.Find(adj) == root {
				newBlock := esploraBlocco(p, adj, seen)
				otherBlocks = append(otherBlocks, newBlock)
			}
		}
	}

    totalLostIntensity := 0

	for _, block := range otherBlocks {
		var newRoot piastrella
		blockIntensity := 0
        maxRank := 0

		for t := range block {
			blockIntensity += p.tiles[t].intensity
			if p.tiles[t].rank > maxRank {
                maxRank = p.tiles[t].rank
            }
            if newRoot == (piastrella{}) {
                newRoot = t
            }
		}

		for t := range block {
			p.tiles[t].parent = newRoot
		}

		p.tiles[newRoot].blockIntensity = blockIntensity
        p.tiles[newRoot].rank = maxRank
		totalLostIntensity += blockIntensity

        if p.tiles[root].blockIntensity < originalBlockIntensity - totalLostIntensity {
            p.tiles[root].blockIntensity = originalBlockIntensity - totalLostIntensity
        } else {
            p.tiles[root].blockIntensity -= totalLostIntensity
        }

        if p.tiles[root].blockIntensity == 0 {
            p.tiles[root].rank = 0
        }
	}

	p.tiles[root].blockIntensity -= originalBlockIntensity - totalLostIntensity
    p.tiles[root].rank = 0
}
