package main

func esploraVicini(p piano, x, y int, seen map[piastrella]bool, sum *int) {
    tile := piastrella{x: x, y: y}
    adiacenti := getAdiacenti(x, y)
    for _, adj := range adiacenti {
        if _, exists := p.tiles[adj]; exists {
            seen[piastrella{x:x,y:y}] = true
            if p.tiles[adj].color == p.tiles[tile].color {
                *sum += p.tiles[adj].intensity
                esploraVicini(p, adj.x, adj.y, seen, sum)
            }
        }
    }
}

func bloccoOmog(p piano, x, y int) {
    somma := 0
    var sum *int = &somma
    seen := make(map[piastrella]bool)
    seen[piastrella{x: x, y: y}] = true
    esploraVicini(p, x, y, seen, sum)
}
