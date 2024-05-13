package tiles

import (
	"fmt"
	"strconv"
)

func stato(p piano, x, y int) (string, int) {
	if _, exists := p.tiles[piastrella{x, y}]; exists {
		fmt.Println(p.tiles[piastrella{x, y}][0], p.tiles[piastrella{x, y}][1])
		intensity, _ := strconv.Atoi(p.tiles[piastrella{x, y}][1])
		return p.tiles[piastrella{x, y}][0], intensity
	} else {
		return "null", 0
	}
}
