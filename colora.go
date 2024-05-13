package tiles

func colora(p piano, x, y int, alpha string) {
    var tile piastrella 
    tile.x = x
    tile.y = y
    p.tiles[tile] = alpha
} 
