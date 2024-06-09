//Elia Cortesi 01911A

package main

func colora(p piano, x, y int, alpha string, i int) {
    var tile = piastrella{x: x, y: y}
    p.Add(tile, alpha, i)
    var adiacenti = getAdiacenti(x, y)
    for _, adj := range adiacenti {
        if _, exists := p.tiles[adj]; exists {
            p.Union(tile, adj)
        }
    }
}
