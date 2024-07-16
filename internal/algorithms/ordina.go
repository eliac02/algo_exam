package algorithms

import models "tiles/internal/models"

// Ordina sorts the list of rules depending on their usage
//
// @param p The system tiles-rules
func Ordina(p models.Piano) {
	bucketSort(p.Rules)
}
