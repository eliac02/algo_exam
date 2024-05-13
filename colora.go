package tiles

import "strconv"

func colora(p piano, x, y int, alpha string, beta int) {
	var tile piastrella
	tile.x = x
	tile.y = y
	p.tiles[tile] = [2]string{alpha, strconv.Itoa(beta)} // beta is the intensity of the color
}
