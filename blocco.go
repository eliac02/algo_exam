package tiles

import (
	"fmt"
)

func blocco(p piano, x, y int) {
	if _, exists := p.tiles[piastrella{x, y}]; exists {
		// do something
	} else {
		fmt.Println(0)
	}
}
