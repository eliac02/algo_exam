package algorithms

import (
	"strconv"
	"strings"
	models "tiles/internal/models"
)

// Regola inserts the rule in the list of rules of the system
//
// @param p The system tiles-rules
// @param r The command given to the program
func Regola(p models.Piano, r string) {
	rulesSplitted := strings.Split(r, " ")

	// create the rule and add it to the system
	var newRule models.Rule
	newRule.Color = rulesSplitted[0]
	newRule.Ruleset = make([]models.Ruleset, 0)
	for i := 1; i <= (len(rulesSplitted) - 2); i += 2 {
		num, err := strconv.Atoi(rulesSplitted[i])
		if err != nil {
			return
		}
		newRule.Ruleset = append(newRule.Ruleset, models.Ruleset{
			Color: rulesSplitted[i+1],
			Count: num,
		})
	}
	newRule.Usage = 0
	*p.Rules = append(*p.Rules, newRule)
}
