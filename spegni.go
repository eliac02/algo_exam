package main

func spegni(p piano, x, y int) {
    tile := piastrella{x: x, y: y}
    if _, exists := p.tiles[tile]; !exists {
        return
    }
    
    root := p.Find(tile)
    p.tiles[root].blockIntensity -= p.tiles[tile].intensity
    delete(p.tiles, tile)
    
    adiacenti := getAdiacenti(x, y)
    var otherBlocks []map[piastrella]*properties
    seen := make(map[piastrella]bool)

    for _, adj := range adiacenti {
        if _, exists := p.tiles[adj]; exists && !seen[adj] {
            if p.Find(adj) == root {
                newBlock := esploraBlocco(p, adj, seen)
                otherBlocks = append(otherBlocks, newBlock)
            }
        }
    }

    for _, block := range otherBlocks {
        var newRoot piastrella
        for t := range block {
            newRoot = t
            break
        }
        for t := range block {
            p.tiles[t].parent = newRoot
        }

        sumIntensities := 0
        for _, props := range block {
            sumIntensities += props.intensity
        }
        p.tiles[newRoot].blockIntensity = sumIntensities
    }
}
