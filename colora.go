package main

func colora(p piano, x, y int, alpha string, i int) {
    var tile = piastrella{x: x, y: y}
    p.Add(tile, alpha, i)
    var adiacenti = [8]piastrella{{x, y+1}, {x-1, y+1}, {x-1, y}, {x-1, y-1}, {x, y-1}, {x+1, y-1}, {x+1, y}, {x+1, y+1}}
    for _, adj := range adiacenti {
        if _, exists := p.tiles[adj]; exists {
            p.Union(tile, adj)
        }
    }
}
