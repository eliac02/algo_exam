package tiles

func spegni(p piano, x, y int) {
    var tile piastrella
    tile.x = x
    tile.y = y
    delete(p.tiles, tile)
}
