package algorithms

import "tiles/internal/models"

// MakeSet creates a system tiles-rules
//
// @return piano The system tiles-rules created
func MakeSet() models.Piano {
	// create the system
	Rules := make([]models.Rule, 0)
	return models.Piano{
		Tiles: make(map[models.Piastrella]*models.Properties),
		Rules: &Rules,
	}
}


