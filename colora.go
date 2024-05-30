package tiles

func colora(p piano, x, y int, alpha string, i int) {
    var counter int
    var rads = make(map[piastrella]int)
	var tile piastrella
	tile.x = x
	tile.y = y
	surroundings := [8]piastrella{{x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
	for _, s := range surroundings {
        //posso avere al massimo 4 blocchi diversi circonvicini
		if _, exists := p.tiles[s]; exists {
            rads[find(p, s)]++
        } else {
            counter++
        }
	}
    if counter == 8 {
        //significa che non ci sono piastrella circonvicine
        p.tiles[tile] = properties{color: alpha, intensity: i, parent: &tile, rank: 0} // beta is the intensity of the color
    } else {
        //guardo quante radici diverse ci sono
        
    }
}
