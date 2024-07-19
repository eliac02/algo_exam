package algorithms

import models "tiles/internal/models"

// Ordina sorts the list of rules depending on their usage
//
// @param ui The graphic interface
// @param p The system tiles-rules
func Ordina(ui *models.UI, p models.Piano) {
	bucketSort(ui, p.Rules)
}
