package main

func makeSet() piano {
    rules := make([]rule, 0)
    return piano{
        tiles: make(map[piastrella]*properties),
        rules: &rules,
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
    rootY := p.Find(y)

    if rootX != rootY {
        if p.tiles[rootX].rank > p.tiles[rootY].rank {
            p.tiles[rootY].parent = rootX
            p.tiles[rootX].rank += p.tiles[rootY].rank+1
            p.tiles[rootY].rank = 0
            p.tiles[rootX].blockIntensity += p.tiles[rootY].blockIntensity
            p.tiles[rootY].blockIntensity = p.tiles[rootY].intensity
        } else if p.tiles[rootY].rank > p.tiles[rootX].rank {
            p.tiles[rootX].parent = rootY
            p.tiles[rootY].rank += p.tiles[rootX].rank+1
            p.tiles[rootX].rank = 0
            p.tiles[rootY].blockIntensity += p.tiles[rootX].blockIntensity
            p.tiles[rootX].blockIntensity = p.tiles[rootX].intensity
        } else {
            p.tiles[rootX].parent = rootY
            p.tiles[rootY].rank++
            p.tiles[rootY].blockIntensity += p.tiles[rootX].blockIntensity
        }
    }
}

func (p piano) Add(x piastrella, c string, i int) {
    if _, exists := p.tiles[x]; !exists {
        p.tiles[x] = &properties{color: c, intensity: i, parent: x, rank: 0, blockIntensity: i}
    } else {
        root := p.Find(x)
        p.tiles[x].color = c
        oldInt := p.tiles[x].intensity
        p.tiles[x].intensity = i
        p.tiles[root].blockIntensity += i-oldInt
    }
}
