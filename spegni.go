package main

func spegni(p piano, x, y int) {
    tile := piastrella{x: x, y: y}
    if _, exists := p.tiles[tile]; !exists {
        return
    }
    
    rootP := p.Find(tile)
    p.tiles[rootP].blockIntensity -= p.tiles[tile].intensity
    delete(p.tiles, tile)
    

}
