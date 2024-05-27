package tiles

func propagaBlocco(p piano, x, y int) {
    if _, exists := p.tiles[piastrella{x, y}]; exists {
        // do something
    } else {
        return
    }
}
