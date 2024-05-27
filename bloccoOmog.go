package tiles

import (
	"fmt"
)

func bloccoOmog(p piano, x, y int) {
	if _, exists := p.tiles[piastrella{x, y}]; exists {
		// do something
	} else {
		fmt.Println(0)
	}
}
