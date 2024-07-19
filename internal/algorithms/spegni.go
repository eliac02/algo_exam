package algorithms

import (
	"image/color"
	models "tiles/internal/models"
	utils "tiles/internal/utils"
)

// Spegni removes from the system the tile (x,y)
//
// @param ui The graphic interface
// @param p The system tiles-rules
// @param x y The coordinates of the tile
func Spegni(ui *models.UI, p models.Piano, x, y int) {
	tile := models.Piastrella{X: x, Y: y}

	if _, exists := p.Tiles[tile]; !exists {
		return
	}

	root := p.Find(tile)

	// if the tile is the root, change the root
	if tile == root {
		utils.CambiaRadice(p, root)
	}

	// retrieve the new blocks that may have been created
	adiacenti := utils.GetAdiacenti(x, y)
	seen := make(map[models.Piastrella]bool)
	seen[tile] = true
	otherBlocks := make([]map[models.Piastrella]*models.Properties, 0)

	for _, adj := range adiacenti {
		if _, exists := p.Tiles[adj]; exists && !seen[adj] {
			newBlock := utils.TrovaBlocco(p, adj, seen, func(models.Piastrella) bool {
				return true
			})
			otherBlocks = append(otherBlocks, newBlock)
		}
	}

	// fix the information of the new blocks
	if len(otherBlocks) > 0 {
		for _, block := range otherBlocks {
			var newRoot models.Piastrella
			totalIntensity := 0
			maxSize := 0

			for t := range block {
				totalIntensity += p.Tiles[t].Intensity
				maxSize++
				if newRoot == (models.Piastrella{}) {
					newRoot = t
				}
				p.Tiles[t].Parent = newRoot
				if p.Tiles[t].BlockIntensity > p.Tiles[t].Intensity {
					p.Tiles[t].BlockIntensity = p.Tiles[t].Intensity
				}
				if p.Tiles[t].Size > 0 {
					p.Tiles[t].Size = 0
				}
			}

			p.Tiles[newRoot].BlockIntensity = totalIntensity
			p.Tiles[newRoot].Size = maxSize
		}
	}

	// ddelete the tile from the system
	delete(p.Tiles, tile)

    ui.UpdateCell(x, 14-y, color.Transparent)
}
