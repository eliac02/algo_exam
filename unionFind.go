package main

func makeSet() piano {
    return piano{
        tiles: make(map[piastrella]*properties),
        rules: make([]rule, 0),
    }
}

func (p piano) Find(x piastrella) piastrella {
    if p.tiles[x].parent != x {
        p.tiles[x].parent = p.Find(p.tiles[x].parent)
    }
    return p.tiles[x].parent
}

func (p piano) Union(x, y piastrella) {
    rootX := p.Find(x)
    rootY := p.Find(x)

    if rootX != rootY {
        if p.tiles[rootX].rank > p.tiles[rootY].rank {
            p.tiles[rootY].parent = rootX
            p.tiles[rootX].blockIntensity += p.tiles[rootY].blockIntensity
        } else if p.tiles[rootY].rank > p.tiles[rootX].rank {
            p.tiles[rootX].parent = rootY
            p.tiles[rootY].blockIntensity += p.tiles[rootX].blockIntensity
        } else {
            p.tiles[rootY].parent = rootX
            p.tiles[rootX].rank++
            p.tiles[rootX].blockIntensity += p.tiles[rootY].blockIntensity
        }
    }
}

func (p piano) Add(x piastrella, c string, i int) {
    if _, exists := p.tiles[x]; !exists {
        p.tiles[x] = &properties{color: c, intensity: i, parent: x, rank: 0, blockIntensity: i}
    }
}
