package tiles

func colora(p piano, x, y int, alpha string, i int) {
	var tile piastrella
	tile.x = x
	tile.y = y
	surroundings := [8]piastrella{{x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
	for _, s := range surroundings {
		if _, exists := p.tiles[s]; exists {
			p.tiles[tile] = properties{color: alpha, intensity: i} // beta is the intensity of the color
		}
	}
}
