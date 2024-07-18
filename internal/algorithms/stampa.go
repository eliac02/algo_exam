package algorithms

import (
	"fmt"
	models "tiles/internal/models"
)

// Stampa prints the list of the rules of the system
//
// @param p The system tiles-rules
func Stampa(p models.Piano) {
	// print the rules
	fmt.Println("(")
	for _, rule := range *p.Rules {
		fmt.Printf("%s:", rule.Color)
		for _, rs := range rule.Ruleset {
			fmt.Printf(" %d %s", rs.Count, rs.Color)
		}
		fmt.Println()
	}
	fmt.Printf(")\n")
}
