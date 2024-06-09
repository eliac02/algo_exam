//Elia Cortesi 01911A

package main

import (
	"strconv"
	"strings"
)

// regola inserts the rule in the list of rules of the system
//
// @param p The system tiles-rules
// @param r The command given to the program
func regola(p piano, r string) {
	rulesSplitted := strings.Split(r, " ")

    //create the rule and add it to the system
	var newRule rule
	newRule.raw = r
	newRule.color = rulesSplitted[1]
	newRule.ruleset = make([]ruleset, 0)
	for i := 2; i <= (len(rulesSplitted) - 2); i += 2 {
		num, err := strconv.Atoi(rulesSplitted[i])
		if err != nil {
			return
		}
        newRule.ruleset = append(newRule.ruleset, ruleset{
            color: rulesSplitted[i+1],
            count: num,
        })
	}
	newRule.usage = 0
    *p.rules = append(*p.rules, newRule)
}
