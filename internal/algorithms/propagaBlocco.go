package algorithms

import models "tiles/internal/models"

// PropagaBlocco applies the first rule available to each of the tiles of the block that the tile (x,y) belongs to
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
func PropagaBlocco(p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}

	// retrieve the block
	root := p.Find(tile)
	seen := make(map[models.Piastrella]bool)
	block := trovaBlocco(p, root, seen, func(models.Piastrella) bool {
		return true
	})

	originalBlock := make(map[models.Piastrella]*models.Properties)
	for t, props := range block {
		originalBlock[t] = &models.Properties{
			Color:          props.Color,
			Intensity:      props.Intensity,
			Parent:         props.Parent,
			Size:           props.Size,
			BlockIntensity: props.BlockIntensity,
		}
	}

	// check for every tile if a rule can be applied
	for t := range block {
		colorCount := make(map[string]int)
		adiacenti := getAdiacenti(t.X, t.Y)
		for _, adj := range adiacenti {
			if props, exists := originalBlock[adj]; exists {
				colorCount[props.Color]++
			}
		}

		for index, reg := range *p.Rules {
			if ruleOk(reg, colorCount) {
				p.Tiles[t].Color = reg.Color
				(*p.Rules)[index].Usage++
				break
			}
		}

	}
}
