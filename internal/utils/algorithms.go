package utils

import (
	"strings"
	models "tiles/internal/models"
)

// GetAdiacenti returns the list of the tiles adjacents to the tile (x,y)
//
// @param x y The coordinates of the tile
// @return [8]piastrella The list of the tiles adjacents to the tile (x,y)
func GetAdiacenti(x, y int) [8]models.Piastrella {
	// retrieve the adjacents tiles
	adjs := [8]models.Piastrella{
		{X: x, Y: y + 1}, {X: x - 1, Y: y + 1}, {X: x - 1, Y: y}, {X: x - 1, Y: y - 1}, {X: x, Y: y - 1}, {X: x + 1, Y: y - 1}, {X: x + 1, Y: y}, {X: x + 1, Y: y + 1},
	}
	return adjs
}

// RuleOk checks if a rule can be applied to a tile
//
// @param reg The rule to check the validity of
// @colorCount A map that contains the color of the tiles adjacents to a tile
// @return bool True if the rule can be applied, False otherwise
func RuleOk(reg models.Rule, colorCount map[string]int) bool {
	// check if the rule can be applied
	for _, rs := range reg.Ruleset {
		if colorCount[rs.Color] < rs.Count {
			return false
		}
	}
	return true
}

// TrovaBlocco finds all the tiles belonging to the block that the tile start belongs to
//
// @param p The system tiles-rules
// @param start The tile which start the search from
// @param filtro Lambda function that checks a given requirement
// @return map[piastrella]*properties The block that the tile start belongs to
func TrovaBlocco(p models.Piano, start models.Piastrella, seen map[models.Piastrella]bool, filtro func(p models.Piastrella) bool) map[models.Piastrella]*models.Properties {
	block := make(map[models.Piastrella]*models.Properties)
	pila := []models.Piastrella{start}

	// use a DFS to retrieve the block's tiles
	for len(pila) > 0 {
		current := pila[len(pila)-1]
		pila = pila[:len(pila)-1]

		if seen[current] {
			continue
		}

		seen[current] = true
		block[current] = p.Tiles[current]

		adiacenti := GetAdiacenti(current.X, current.Y)
		for _, adj := range adiacenti {
			if _, exists := p.Tiles[adj]; exists && !seen[adj] {
				// use a lambda function as a filter
				if filtro(adj) {
					pila = append(pila, adj)
				}
			}
		}
	}
	return block
}

// CambiaRadice change the root of a block
//
// @param p The system tiles-rules
// @param root The root that needs to be changes
func CambiaRadice(p models.Piano, root models.Piastrella) {
	adiacenti := GetAdiacenti(root.X, root.Y)
	seen := make(map[models.Piastrella]bool)
	block := TrovaBlocco(p, root, seen, func(models.Piastrella) bool {
		return true
	})

	// change the root to the first adjacent tile found
	for _, adj := range adiacenti {
		if _, exists := p.Tiles[adj]; exists {
			p.Tiles[root].BlockIntensity -= p.Tiles[root].Intensity
			p.Tiles[root].Size--
			p.Tiles[adj].BlockIntensity = p.Tiles[root].BlockIntensity
			p.Tiles[adj].Size = p.Tiles[root].Size
			p.Tiles[adj].Parent = adj
			for t := range block {
				p.Tiles[t].Parent = adj
			}
			return
		}
	}
}

// VerificaPista checks if the sequence of directions is valid
//
// @param p The system tiles-rules
// @param x y The coordinates of the tile
// @param directions The list of directions
// @return bool True if the sequence is valid, False otherwise
func VerificaPista(p models.Piano, x, y int, s string) (bool, []models.Piastrella) {
	sequence := make([]models.Piastrella, 0)
	// assign to each direction an index
	corresponding := map[string]int{
		"NN": 0,
		"NO": 1,
		"OO": 2,
		"SO": 3,
		"SS": 4,
		"SE": 5,
		"EE": 6,
		"NE": 7,
	}

	directions := strings.Split(s, ",")

	// check one-by-one if all the directions are valid
	for _, dir := range directions {
		adiacenti := GetAdiacenti(x, y)
		tile := adiacenti[corresponding[dir]]
		if _, exists := p.Tiles[tile]; !exists {
			return false, []models.Piastrella{}
		}
		sequence = append(sequence, tile)
		x, y = tile.X, tile.Y
	}
	return true, sequence
}

// CamminoMinimo calculate the shortest path from start to end
//
// @param p The system tiles-rules
// @param start The starting tile
// @param end The ending tile
// @return int The lenght of the shortest path between @start and @end
func CamminoMinimo(p models.Piano, start, end models.Piastrella) int {
	coda := []models.Piastrella{start}
	distanze := make(map[models.Piastrella]int)
	distanze[start] = 1

	// use a BFS to calculate the shortest path
	for len(coda) > 0 {
		current := coda[0]
		coda = coda[1:]

		if current == end {
			return distanze[current]
		}

		adiacenti := GetAdiacenti(current.X, current.Y)
		for _, adj := range adiacenti {
			if _, exists := p.Tiles[adj]; exists {
				if _, visited := distanze[adj]; !visited {
					coda = append(coda, adj)
					distanze[adj] = distanze[current] + 1
				}
			}
		}
	}

	return 0
}
