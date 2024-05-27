package tiles

func colora(p piano, x, y int, alpha string, i int) {
	var tile piastrella
	tile.x = x
	tile.y = y
	p.tiles[tile] = properties{color: alpha, intensity: i} // beta is the intensity of the color
}
